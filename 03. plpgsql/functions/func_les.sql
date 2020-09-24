-- ищет по фамилии id человека
CREATE OR REPLACE FUNCTION get_user_id(surname varchar) RETURNS int
AS $$
DECLARE
user_id int;
BEGIN
	SELECT people.id INTO user_id
	FROM people
	WHERE people.surname = get_user_id.surname;
	RETURN user_id;
END
$$ LANGUAGE plpgsql;

-- добавление STRICT будет возвращать ошибку, если не найдено
CREATE OR REPLACE FUNCTION get_user_id(surname varchar) RETURNS int
AS $$
DECLARE
user_id int;
BEGIN
	SELECT people.id INTO STRICT user_id
	FROM people
	WHERE people.surname = get_user_id.surname;
	RETURN user_id;
END
$$ LANGUAGE plpgsql;

-- для STRICT мы можем написать конструкцию EXCEPTION, чтобы
-- обрабатывать ошибки, в данном случае мы перехватываем NO_DATA_FOUND,
-- чтоб вернуть свою ошибку
CREATE OR REPLACE FUNCTION get_user_id(surname varchar) RETURNS int
AS $$
DECLARE
user_id int;
BEGIN
	SELECT people.id INTO STRICT user_id
	FROM people
	WHERE people.surname = get_user_id.surname;
	RETURN user_id;
	EXCEPTION
		WHEN NO_DATA_FOUND THEN
			RAISE EXCEPTION 'people % not found', surname;
END
$$ LANGUAGE plpgsql;

-- если запустить так
SELECT get_user_id('orlova')

-- будет ошибка, что слишком много строк, в INTO можно положить только одну строку!

-- можем поменять так
CREATE OR REPLACE FUNCTION get_user_id(surname varchar) RETURNS int
AS $$
DECLARE
user_id int;
BEGIN
	SELECT people.id INTO STRICT user_id
	FROM people
	WHERE people.surname = get_user_id.surname;
	RETURN user_id;
	EXCEPTION
		WHEN NO_DATA_FOUND THEN
			RAISE EXCEPTION 'people % not found', surname;
		WHEN TOO_MANY_ROWS THEN
      RAISE EXCEPTION 'people % not unique', surname;
END
$$ LANGUAGE plpgsql;

-- и получать ошибку а-ля 'people orlova not unique'

-- Если ошибка не приемлимый вариант, то можем добавить LIMIT в запрос,
-- но тогда будем получать только одного случайно (а вдруг их там много?)

-- функция, которая кладёт результат в переменную, которая включает в себя
-- все атрибуты указанной таблицы, обращение к атрибутом по этой таблице через
-- точку
CREATE OR REPLACE FUNCTION get_user_info(id int) RETURNS varchar
AS $$
DECLARE
t_people people%ROWTYPE;
BEGIN
	SELECT * INTO t_people
	FROM people
	WHERE people.id = get_user_info.id;
	
	RETURN t_people.name || ' ' || t_people.surname;
END
$$ LANGUAGE plpgsql;

-- функция также может возвращать QUERY
CREATE OR REPLACE FUNCTION get_query_info(id int) RETURNS SETOF people AS 
$BODY$
BEGIN
	RETURN QUERY 
		SELECT * 
		FROM people
		WHERE people.id = get_query_info.id;
END
$BODY$
LANGUAGE plpgsql;

-- Её можно запустить так:
SELECT * FROM get_query_info(6)

-- Можно и так:
SELECT get_query_info(6)

-- но получите не совсем ожидаемый результат :)