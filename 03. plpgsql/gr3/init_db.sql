CREATE TABLE staff (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	surname varchar(255) NOT NULL,
	position varchar(255) NOT NULL,
	first_day DATE NOT NULL,
	experience_before_month INTEGER NOT NULL DEFAULT 0,
	birth_day DATE NOT NULL,
	sex varchar(1) NOT NULL
);

CREATE TABLE patient (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	surname varchar(255) NOT NULL,
	birth_day DATE NOT NULL,
	sex varchar(1) NOT NULL
);


CREATE TABLE consultation ( 
	id SERIAL PRIMARY KEY,
	patient_id INTEGER NOT NULL REFERENCES patient(id),
	staff_id INTEGER NOT NULL REFERENCES staff(id),
	arrive_date DATE NOT NULL DEFAULT now(),
	description TEXT
);


CREATE TABLE schedule (
	id SERIAL PRIMARY KEY,
	staff_id INTEGER NOT NULL REFERENCES staff(id),
	work_start TIMESTAMP NOT NULL,
	work_end TIMESTAMP NOT NULL
);

INSERT INTO staff (name, surname, position, first_day, experience_before_month, birth_day, sex) VALUES 
('Иван', 'Иванов', 'Терапевт', '11.01.2013', 23, '01.15.1970', 'М'),
('Антон', 'Сидоров', 'Хирург', '07.14.2008', 120, '03.08.1963', 'М');

INSERT INTO patient (name, surname, birth_day, sex) VALUES 
('Александр', 'Александров', '11.23.1987', 'М'),
('Жанна', 'Мякишева', '08.12.1993', 'Ж');

INSERT INTO consultation (patient_id, staff_id, arrive_date, description) VALUES 
(1, 1, now(), 'Пациент пришёл на прием');

INSERT INTO schedule (staff_id, work_start, work_end) VALUES
(1, '10.10.2020 11:00', '10.10.2020 16:00'),
(1, '10.11.2020 11:00', '10.11.2020 16:00'),
(2, '10.11.2020 10:00', '10.11.2020 14:00');

CREATE OR REPLACE FUNCTION is_staff_working(t Timestamp, id INTEGER) Returns Boolean
AS $$
DECLARE
  schedule_id INTEGER;
BEGIN
  Select schedule.id INTO schedule_id
  From schedule
  where staff_id = is_staff_working.id
  and work_start <= t and work_end >= t;

  If schedule_id is NULL THEN
    RETURN FALSE;
  END IF;

  RETURN TRUE;
END
$$ Language plpgsql;

-- SELECT is_staff_working('10.11.2020 13:00', 1)

-- вывести персонал с указанной позицией, работающий в заданное время
-- вывод просто в сообщения процедуры
CREATE OR REPLACE PROCEDURE find_working_doctor(
	t timestamp ,
	v varchar)
LANGUAGE 'plpgsql'
AS $$
DECLARE
	s staff%ROWTYPE;
BEGIN
	For s IN
		Select staff.* From staff
		Inner Join schedule on schedule.staff_id = staff.id
		Where position = v
		and work_start <= t and work_end >= t
	LOOP
		RAISE NOTICE 'name: %, surname: %', s.name,s.surname;
	END LOOP;
END;
$$;

-- вывести персонал с указанной позицией, работающий в заданное время
CREATE OR REPLACE FUNCTION find_working_doctor_f(t TIMESTAMP, v VARCHAR) RETURNS table(name varchar, surname varchar)
AS $$
BEGIN
  RETURN QUERY 
    Select staff.name, staff.surname From staff
    Inner Join schedule on schedule.staff_id = staff.id
    Where position = v
    and work_start <= t and work_end >= t;
END
$$ Language plpgsql;

-- вывести персонал с указанной позицией, работающий в заданное время (С SETOF)
CREATE OR REPLACE FUNCTION find_working_doctor_f(t TIMESTAMP, v VARCHAR) RETURNS SETOF staff
AS $$
BEGIN
  RETURN QUERY 
    Select staff.* From staff
    Inner Join schedule on schedule.staff_id = staff.id
    Where position = v
    and work_start <= t and work_end >= t;
END
$$ Language plpgsql;

/* #### LESSON 24.10 #### */
CREATE TABLE staff_audit (
  -- системные атрибуты
  operation varchar(1) NOT NULL, -- Либо INSERT (I), UPDATE (U), DELETE (D); TG_OP
  created_at TIMESTAMP DEFAULT now() NOT NULL, -- Во сколько была создана строка в таблице staff audit
  -- атрибуты из таблицы staff
  staff_id INTEGER NOT NULL, -- тут осознанно нет REFERENCES, потому что в случае удаления в таблице staff ссылаться будет не на что 
  name varchar(255) NOT NULL,
  surname varchar(255) NOT NULL,
  position varchar(255) NOT NULL,
  first_day DATE NOT NULL,
  experience_before_month INTEGER NOT NULL DEFAULT 0,
  birth_day DATE NOT NULL,
  sex varchar(1) NOT NULL
);

CREATE OR REPLACE FUNCTION staff_audit_tg() RETURNS TRIGGER AS $$
  BEGIN
    -- если триггер был вызван оператором INSERT
    IF (TG_OP = 'INSERT') THEN
      RAISE NOTICE 'RUN INSERT in staff_audit_tg';
      INSERT INTO staff_audit SELECT 'I', now(), NEW.*;
    ELSIF (TG_OP = 'UPDATE') THEN
      RAISE NOTICE 'RUN UPDATE in staff_audit_tg';
      INSERT INTO staff_audit SELECT 'U', now(), NEW.*;
    ELSIF (TG_OP = 'DELETE') THEN
      RAISE NOTICE 'RUN DELETE in staff_audit_tg';
      INSERT INTO staff_audit SELECT 'D', now(), OLD.*;
    END IF;
    RETURN NULL; -- нам не нужно ничего возвращать, т.к. функция будет использоваться в AFTER tg
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER staff_audit
  AFTER INSERT OR UPDATE OR DELETE ON staff
  FOR EACH ROW EXECUTE FUNCTION staff_audit_tg();
  
INSERT INTO staff (name, surname, position, first_day, experience_before_month, birth_day, sex) VALUES 
('Пётр', 'Петров', 'Окулист', '11.01.2008', 126, '01.15.1961', 'М');

UPDATE staff
SET surname = 'Николаев'
WHERE id = 4;

DELETE FROM staff WHERE id = 3;

SELECT * FROM staff;

SELECT * FROM staff_audit;

SELECT * FROM patient;

CREATE OR REPLACE VIEW patient_consultation_view AS
  SELECT patient.*, consultation.staff_id, consultation.arrive_date, consultation.description
  FROM patient
  INNER JOIN consultation ON consultation.patient_id = patient.id;

-- изменяем тип arrive_date на TIMESTAMP
ALTER TABLE consultation 
  ALTER COLUMN arrive_date TYPE TIMESTAMP;
  
SELECT * 
FROM patient_consultation_view;

INSERT INTO patient_consultation_view(id, staff_id, arrive_date, description) VALUES
(3, 2, now(), 'Description...')


UPDATE patient_consultation_view
SET arrive_date = now(), description = 'Changes were made...'
WHERE consultation_id = 5;

DELETE FROM patient_consultation_view
WHERE consultation_id = 3;

CREATE OR REPLACE VIEW patient_consultation_view AS
SELECT patient.*, consultation.id as consultation_id, consultation.staff_id, consultation.arrive_date, consultation.description
FROM patient
INNER JOIN consultation ON consultation.patient_id = patient.id;

DROP VIEW patient_consultation_view

CREATE OR REPLACE FUNCTION update_patient_consultation_view() RETURNS TRIGGER AS $$
BEGIN
	IF (TG_OP = 'INSERT') THEN
		RAISE NOTICE 'RUN INSERT in staff_audit_tg';
		INSERT INTO consultation(patient_id, staff_id, arrive_date, description) VALUES
		(NEW.id, NEW.staff_id, NEW.arrive_date, NEW.description);
		RETURN NEW;
	ELSIF (TG_OP = 'UPDATE') THEN
		RAISE NOTICE 'RUN UPDATE in staff_audit_tg';
		UPDATE consultation
		SET arrive_date = NEW.arrive_date, description = NEW.description
		WHERE id = OLD.consultation_id;
		RETURN NEW;
	ELSIF (TG_OP = 'DELETE') THEN
		RAISE NOTICE 'RUN DELETE in staff_audit_tg';
		DELETE FROM consultation
		Where id = OLD.consultation_id;
		RETURN OLD;
	END IF;
END
$$ LANGUAGE plpgsql;


CREATE TRIGGER patient_consultation_view_tg
INSTEAD OF INSERT OR UPDATE OR DELETE
ON patient_consultation_view
FOR EACH ROW EXECUTE FUNCTION update_patient_consultation_view();

CREATE TRIGGER not_allow_change_name_tg
BEFORE UPDATE — or AFTER
ON staff
FOR EACH ROW EXECUTE FUNCTION not_allow_change_name();

UPDATE staff
SET name = 'Doe'
WHERE id = 1;