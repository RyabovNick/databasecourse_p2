-- Триггер, который проверяет, если rec_user.username старый и новый не совпадают, то вывести ошибку, что изменять username нельзя
CREATE OR REPLACE TRIGGER check_username
BEFORE UPDATE ON rec_user
FOR EACH ROW
-- если нужно объявить какие-то переменные, то в DECLARE
BEGIN
    IF :NEW.username <> :OLD.username THEN
        -- Например, если мы не хотим возвращать ошибку, а просто запретить изменять, то можем сделат
        -- :NEW.username := :OLD.username
        raise_application_error(-20200, 'Username can not be changed');
    END IF;
END;

SELECT sysdate FROM rec_user;

UPDATE rec_user
SET username = 'username22'
WHERE id = 21; -- этот update вернёт нашу ошибку (raise_application_error(-20200, 'Username can not be changed');)

-- Сделать аудит таблицы rec_ingredient
-- Записываем каждое действие в таблицу (и ставим что это было: DELETE (D), INSERT (I), UPDATE (U), также пишем дату, когда произошло)
CREATE TABLE rec_ingredient_audit (
    id NUMBER GENERATED ALWAYS as IDENTITY(START with 1 INCREMENT by 1),
    rec_ingredient_id NUMBER NOT NULL, -- это ID из таблицы rec_ingredient, но мы не делаем внешний ключ. 
    name varchar2(255) NOT NULL,
    calorie real,
    fat real,
    protein real, 
    carbohydrate real,
    amount INTEGER,
    action varchar2(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);

SELECT * FROM rec_ingredient;

CREATE OR REPLACE TRIGGER rec_ingredient_audit_tg
AFTER INSERT OR UPDATE OR DELETE 
ON rec_ingredient
FOR EACH ROW
DECLARE
    action varchar2(1);
BEGIN
    -- в зависимости от того, на какую команду сработал триггер, определяем значение атрибута action
    IF inserting THEN 
        action := 'I'; 

        INSERT INTO rec_ingredient_audit (rec_ingredient_id, name, calorie, fat, protein, carbohydrate, amount, action)
        VALUES (:NEW.id, :NEW.name, :NEW.calorie, :NEW.fat, :NEW.protein, :NEW.carbohydrate, :NEW.amount, action);
    END IF;

    IF updating THEN 
        action := 'U'; 

        INSERT INTO rec_ingredient_audit (rec_ingredient_id, name, calorie, fat, protein, carbohydrate, amount, action)
        VALUES (:NEW.id, :NEW.name, :NEW.calorie, :NEW.fat, :NEW.protein, :NEW.carbohydrate, :NEW.amount, action);
    END IF;

    IF deleting THEN 
        action := 'D';

        -- в принципе в delete можно и не сохранять значения (их можно найти в предыдущем I or U действии, но на всякий случай мы добавили)
        INSERT INTO rec_ingredient_audit (rec_ingredient_id, name, calorie, fat, protein, carbohydrate, amount, action)
        VALUES (:OLD.id, :OLD.name, :OLD.calorie, :OLD.fat, :OLD.protein, :OLD.carbohydrate, :OLD.amount, action);
    END IF;
END; 

SELECT * FROM rec_ingredient;
SELECT * FROM rec_ingredient_audit;

-- создаём, должна появится соответсвующая запись в таблице rec_ingredient
INSERT INTO rec_ingredient (name, calorie, fat, protein, carbohydrate, amount) VALUES ('Test1', 89, 0.3, 1.1, 23, 100);

UPDATE rec_ingredient
SET calorie = 26, fat = 0.8, protein = 1.4
WHERE id = 21;

DELETE FROM rec_ingredient WHERE id = 21;