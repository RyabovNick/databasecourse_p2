CREATE TABLE rec_user (
    id NUMBER GENERATED ALWAYS as IDENTITY(START with 1 INCREMENT by 1),
    username varchar2(255) NOT NULL UNIQUE,
    email varchar2(255) NOT NULL UNIQUE,
    PRIMARY KEY(id)
);

INSERT INTO rec_user (username, email) VALUES ('username1', 'email1@mail.ru');
INSERT INTO rec_user (username, email) VALUES ('username2', 'email2@mail.ru');

SELECT * FROM rec_user

CREATE TABLE rec_recipe (
    id NUMBER GENERATED ALWAYS as IDENTITY(START with 1 INCREMENT by 1),
    user_id NUMBER NOT NULL REFERENCES rec_user(id),
    title varchar2(255) NOT NULL,
    description varchar2(1000),
    total_time INTEGER,
    people_time INTEGER,
    PRIMARY KEY(id)
);

INSERT INTO rec_recipe (user_id, title, total_time, people_time) VALUES (1, 'Оладушки', 30, 20);
INSERT INTO rec_recipe (user_id, title, total_time, people_time) VALUES (21, 'Борщ', 120, 60);

CREATE TABLE rec_ingredient (
    id NUMBER GENERATED ALWAYS as IDENTITY(START with 1 INCREMENT by 1),
    name varchar2(255) NOT NULL,
    calorie real,
    fat real,
    protein real, 
    carbohydrate real,
    amount INTEGER,
    PRIMARY KEY(id)
);

INSERT INTO rec_ingredient (name, calorie, fat, protein, carbohydrate, amount) VALUES ('Банан', 89, 0.3, 1.1, 23, 100);
INSERT INTO rec_ingredient (name, calorie, fat, protein, carbohydrate, amount) VALUES ('Свекла', 43, 0.2, 1.6, 10, 100);

CREATE TABLE rec_r_i (
    recipe_id NUMBER NOT NULL REFERENCES rec_recipe(id),
    ingredient_id NUMBER NOT NULL REFERENCES rec_ingredient(id),
    amount INTEGER, -- 10 , 150 , 200 , 2 , 3 
    quantity varchar2(255), -- штук, гр, ч.л. ст.л, стакан.....
    PRIMARY KEY (recipe_id, ingredient_id)
);

INSERT INTO rec_r_i (recipe_id, ingredient_id, amount, quantity) VALUES (1, 1, 200, 'гр.');
INSERT INTO rec_r_i (recipe_id, ingredient_id, amount, quantity) VALUES (2, 2, 300, 'гр.');

SELECT *
FROM rec_recipe
INNER JOIN rec_r_i ON rec_recipe.id = rec_r_i.recipe_id
INNER JOIN rec_ingredient ON rec_ingredient.id = rec_r_i.ingredient_id

SELECT * FROM rec_user

-- 1. SELECT INTO - когда 1 строка
-- 2. cursor

-- переменные с явным типом
DECLARE
    username varchar2(255); -- переменная, в которую будет класть на каждой итерации результат из курсора
    email varchar2(255);

    CURSOR all_users IS
        SELECT username, email FROM rec_user;
BEGIN
    -- открываем курсор, он указывает на 1ую строку
    OPEN all_users;
    -- делаем цикл, в котором будем перебирать все строки запроса через курсор
    LOOP
        -- кладём строку, на которую указывает курсор в указанную переменную (или переменные)
        FETCH all_users INTO username, email;
        -- выходим из цикла тогда, когда курсор указывает на данные, которых нет (т.е. ничего не находит дальше, никаких данных)
        EXIT WHEN all_users%NOTFOUND;
        dbms_output.put_line(username || ' ' || email);
    END LOOP;
    -- закрываем курсор
    CLOSE all_users;
END

-- использование %TYPE для объявления типа
DECLARE
    -- %TYPE берёт тот тип, который указан в таблице у конкретного атрибута
    username rec_user.username%TYPE; -- переменная, в которую будет класть на каждой итерации результат из курсора
    email rec_user.email%TYPE;

    CURSOR all_users IS
        SELECT username, email FROM rec_user;
BEGIN
    -- открываем курсор, он указывает на 1ую строку
    OPEN all_users;
    -- делаем цикл, в котором будем перебирать все строки запроса через курсор
    LOOP
        -- кладём строку, на которую указывает курсор в указанную переменную (или переменные)
        FETCH all_users INTO username, email;
        -- выходим из цикла тогда, когда курсор указывает на данные, которых нет (т.е. ничего не находит дальше, никаких данных)
        EXIT WHEN all_users%NOTFOUND;
        dbms_output.put_line(username || ' ' || email);
    END LOOP;
    -- закрываем курсор
    CLOSE all_users;
END

-- использование %ROWTYPE для объявления типа на основе таблицы в базе
DECLARE
    -- %ROWTYPE берёт все атрибуты из указанной таблицы с указанными в ней типами
    r_u rec_user%ROWTYPE; -- переменная, в которую будет класть на каждой итерации результат из курсора

    CURSOR all_users IS
        SELECT * FROM rec_user;
BEGIN
    -- открываем курсор, он указывает на 1ую строку
    OPEN all_users;
    -- делаем цикл, в котором будем перебирать все строки запроса через курсор
    LOOP
        -- кладём строку, на которую указывает курсор в указанную переменную (или переменные)
        FETCH all_users INTO r_u;
        -- выходим из цикла тогда, когда курсор указывает на данные, которых нет (т.е. ничего не находит дальше, никаких данных)
        EXIT WHEN all_users%NOTFOUND;
        dbms_output.put_line(r_u.username || ' ' || r_u.email);
    END LOOP;
    -- закрываем курсор
    CLOSE all_users;
END

-- использование %ROWTYPE для объявления типа на основе курсора, объявленного в этом же DECLARE
DECLARE
    CURSOR all_users IS
        SELECT username, email FROM rec_user;

    -- %ROWTYPE берёт все атрибуты из указанного курсора с указанными в нём типами
    r_u all_users%ROWTYPE; -- переменная, в которую будет класть на каждой итерации результат из курсора
BEGIN
    -- открываем курсор, он указывает на 1ую строку
    OPEN all_users;
    -- делаем цикл, в котором будем перебирать все строки запроса через курсор
    LOOP
        -- кладём строку, на которую указывает курсор в указанную переменную (или переменные)
        FETCH all_users INTO r_u;
        -- выходим из цикла тогда, когда курсор указывает на данные, которых нет (т.е. ничего не находит дальше, никаких данных)
        EXIT WHEN all_users%NOTFOUND;
        dbms_output.put_line(r_u.username || ' ' || r_u.email);
    END LOOP;
    -- закрываем курсор
    CLOSE all_users;
END

-- SELECT INTO:
DECLARE
    cnt number;
BEGIN
    SELECT count(*) INTO cnt FROM rec_user;
    dbms_output.put_line(cnt);
END

-- использование FOR для работы с курсором
DECLARE
    CURSOR all_users IS
        SELECT username, email FROM rec_user;
    
    -- явного объявления переменной уже не требуется, это за нас сделает FOR
BEGIN
    -- открывать курсор тоже не требуется, в FOR происходит это автоматически
    -- делаем цикл FOR, в котором будем перебирать все строки запроса через курсор
    FOR r_u IN all_users
    LOOP
        -- FETCH делать не нужно - FOR на каждом шаге это делает за нас
        -- выход из цикла будет осуществлён тогда, когда достигнем последнего
        dbms_output.put_line(r_u.username || ' ' || r_u.email);
    END LOOP;
    -- закрывать не нужно, FOR сделает сам
END


-- процедура
create or replace PROCEDURE all_users AS
    CURSOR all_users IS
        SELECT username, email FROM rec_user;
    
    -- явного объявления переменной уже не требуется, это за нас сделает FOR
BEGIN
    -- открывать курсор тоже не требуется, в FOR происходит это автоматически
    -- делаем цикл FOR, в котором будем перебирать все строки запроса через курсор
    FOR r_u IN all_users
    LOOP
        -- FETCH делать не нужно - FOR на каждом шаге это делает за нас
        -- выход из цикла будет осуществлён тогда, когда достигнем последнего
        dbms_output.put_line(r_u.username || ' ' || r_u.email);
    END LOOP;
    -- закрывать не нужно, FOR сделает сам
END;


create or replace FUNCTION get_user_by_id_func (id NUMBER) RETURN varchar AS 
    user rec_user%ROWTYPE;
BEGIN
    SELECT * 
    INTO user 
    FROM rec_user
    WHERE rec_user.id = get_user_by_id_func.id;
    return user.username;
END;

