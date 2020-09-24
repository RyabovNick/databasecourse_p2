SELECT *
FROM people

-- процедура, которая возводит во 2 степень любое переданное значение
CREATE OR REPLACE PROCEDURE power_two(INOUT x int)
LANGUAGE plpgsql
AS $$
BEGIN
	x := x ^ 2;
END;
$$;

-- вызов процедуры
CALL power_two(5)

-- анонимная функция (не сохраняется)
DO $$
DECLARE var int := 3;
BEGIN
  -- процедура меняет значение переменной var
	CALL power_two(var);
	-- var = 9
  -- RAISE NOTICE - выводит в message сообщение (если pgadmin)
	RAISE NOTICE 'var = %', var;
END;
$$;

-- IF
CREATE OR REPLACE PROCEDURE power_two_if(INOUT x int)
LANGUAGE plpgsql
AS $$
BEGIN
	IF mod(x,2) = 0 THEN
		x:= x ^ 2;
	ELSE
		x := x ^ 3;
	END IF;
END;
$$;

CALL power_two_if(5)

-- CASE
CREATE OR REPLACE PROCEDURE power_two_case(INOUT x int)
LANGUAGE plpgsql
AS $$
BEGIN
	CASE mod(x,2)
		WHEN 0 THEN x := x^2;
		WHEN 1 THEN x := x^3;
	END CASE;
END;
$$;

CALL power_two_case(5)

-- FOR LOOP
CREATE OR REPLACE PROCEDURE loop_proc()
LANGUAGE plpgsql
AS $$
BEGIN
	FOR i IN 1..10 LOOP
		RAISE NOTICE 'its: %', i;
	END LOOP;
END;
$$;

CALL loop_proc()

-- Как использовать FOR LOOP для запросов
CREATE OR REPLACE PROCEDURE get_people()
LANGUAGE plpgsql
AS $$
DECLARE
	p people%ROWTYPE; -- или RECORD, обязательно его, если запрос на несколько таблиц
BEGIN
	RAISE NOTICE 't: %', p.name;
	FOR p IN 
		SELECT * FROM people
	LOOP
		RAISE NOTICE 'name: %, surname: %', p.name, p.surname;
	END LOOP;
END;
$$;

CALL get_people()