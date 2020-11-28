SELECT * FROM client
SELECT * FROM menu
SELECT * FROM order_
SELECT * FROM order_menu

-- функцию, которая ищет по фамилии ID клиента
CREATE OR REPLACE FUNCTION get_client_id(name varchar) RETURNS int
AS $$
	-- объявление переменных
	DECLARE
		client_id int;
	-- исполняемая секция
	BEGIN
		SELECT client.id INTO client_id
		FROM client
		WHERE client.name = get_client_id.name;
		RETURN client_id;
	END;
$$ LANGUAGE plpgsql;

SELECT get_client_id('Вера');

CREATE OR REPLACE FUNCTION get_client_id(name varchar) RETURNS int
AS $$
	-- объявление переменных
	DECLARE
		client_id int;
	-- исполняемая секция
	BEGIN
		SELECT client.id INTO STRICT client_id
		FROM client
		WHERE client.name = get_client_id.name
		LIMIT 1;
		RETURN client_id;
		-- исключения
		EXCEPTION
			WHEN NO_DATA_FOUND THEN
				RAISE EXCEPTION 'client % not found', name;
			WHEN TOO_MANY_ROWS THEN
				RAISE EXCEPTION 'client % not unique', name;
	END;
$$ LANGUAGE plpgsql;

SELECT get_client_id('test4');

-- получить информацию о клиенте
CREATE OR REPLACE FUNCTION get_client_info(id int) RETURNS varchar
AS $$
	-- объявление переменных
	DECLARE
		-- похоже на объект, который содержит все атрибуты таблицы client
		t_client client%ROWTYPE;
		-- t client.id%TYPE; -- ещё вариант объявления типа
	-- исполняемая секция
	BEGIN
		SELECT * INTO t_client
		FROM client
		WHERE client.id = get_client_info.id;
		RETURN t_client.name || ' ' || t_client.phone || ' ' || t_client.address;
		-- исключения
		EXCEPTION
			WHEN NO_DATA_FOUND THEN
				RAISE EXCEPTION 'client % not found', name;
	END;
$$ LANGUAGE plpgsql;

SELECT get_client_info(3);

-- Функция, которая возвращает Query
CREATE OR REPLACE FUNCTION get_client_info_in_query(id int) RETURNS SETOF client
AS $$
	-- исполняемая секция
	BEGIN
		RETURN QUERY 
			SELECT *
			FROM client
			WHERE client.id = get_client_info_in_query.id;
	END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_client_info_in_query(1)


SELECT * FROM menu

-- Каждое изменение меню - пишется в таблицу аудита
-- Может что-то добавиться - запишем, либо измениться - тоже запишем, либо удалиться 
CREATE TABLE menu_audit (
	-- системные атрибуты
	operation varchar(1) NOT NULL, -- INSERT (I), UPDATE (U), DELETE (D) - операция, на которую срабатывает триггер
	tg_created_at TIMESTAMP DEFAULT now() NOT NULL, -- когда была создана строка в этой таблице
	-- атрибуты из таблицы menu
	menu_id INTEGER NOT NULL, -- тут нету REFERENCES, потому что в случае удаления в таблице menu, тут не на что будет ссылаться. Либо из таблицы
		-- menu невозможно удалить, либо при удалении удалить и из menu_audit, что не верно
	name varchar(255) NOT NULL,
	price numeric(8,2) NOT NULL,
	description varchar(3000),
	weight INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION menu_audit_tg() RETURNS TRIGGER AS $$
	BEGIN
		IF (TG_OP = 'INSERT') THEN 
			RAISE NOTICE 'RUN INSERT';
			INSERT INTO menu_audit SELECT 'I', now(), NEW.*;
		ELSIF (TG_OP = 'UPDATE') THEN
			RAISE NOTICE 'RUN UPDATE';
			INSERT INTO menu_audit SELECT 'U', now(), NEW.*;
		ELSIF (TG_OP = 'DELETE') THEN
			RAISE NOTICE 'RUN DELETE';
			INSERT INTO menu_audit SELECT 'D', now(), OLD.*;
		END IF;
		RETURN NULL;
	END;
$$ LANGUAGE plpgsql

-- INSERT INTO table (id, name) VALUES (1, '10')
-- UPDATE table SET name = '2' WHERE id = 1
-- DELETE FROM table WHERE id = 1

CREATE TRIGGER menu_audit
	AFTER -- когда выполняется триггер (BEFORE, INSTEAD OF)
	INSERT OR UPDATE OR DELETE -- на какой(-ие) операторы срабатывает триггер
	ON menu -- на какую таблицу срабатывает триггер
	FOR EACH ROW -- (или может быть не указан, появляется доступ к NEW и OLD)
	-- Для INSERT - есть только NEW.name; 
	-- Для UPDATE - есть и NEW, и OLD. OLD.name - вернёт старый name, до изменения, т.е. текущий в базе. NEW.name - тот, на который меняем
	-- ДЛя DELETE - есть только OLD. Если удаляется, то через OLD.name можно узнать текущее значение в БД, которое будет удалено
	EXECUTE FUNCTION menu_audit_tg();
	
SELECT * FROM menu
SELECT * FROM menu_audit

INSERT INTO menu (name, price, description, weight)
VALUES ('Кофе', 2, null, 300);

UPDATE menu
SET description = 'Кофе!'
WHERE id = 5;

DELETE FROM menu
WHERE id = 5;

UPDATE menu
SET description = 'Классическая пицца'
WHERE id = 1;