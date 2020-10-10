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