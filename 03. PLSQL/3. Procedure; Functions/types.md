# Пользовательские типы данных

Иногда бывает полезным сделать собственные типы, например, когда необходимо, чтобы функция возвращала несколько значений. 

Пользовательский тип - объект, который содержит несколько переменных с любыми типами. (Можете проверить, можно ли использовать пользовательский тип в пользовательском типе и насколько удобно к такому обращаться :D Если можно, то можете узнать насколько такая вложенность может быть ограничена (только зачем?))

В общем, тип необходимо создать. Самый простой вариант

```
CREATE OR REPLACE EDITIONABLE TYPE F_TYPE AS OBJECT
( 
 name_A     VARCHAR2(255),
 name_K      VARCHAR2(255),
 date_F      DATE
);
```

Тип можно использовать, например в функции так:


```
create or replace function f_t_v return F_TYPE is
begin
  return
     F_TYPE('Hello','AH',sysdate);
end;
```

`Select f_t_v from dual`

Запустит функцию.

Посмотрим более сложный вариант.

Необходимо выполнить запрос, положить в переменную тип которой пользовательский несколько значений, вернуть их и запустить функцию потом.


```
create or replace Function type_t
Return F_TYPE 
AS 
c F_TYPE;
BEGIN
    Select k.Name, a.Name, K_A.Finish_D INTO c.name_A, c.name_K, c.date_F
        From Animals a, Keepers k, K_A
        Where a.ID = K_A.A_ID and K_A.K_ID = k.ID and K_A.Finish_D = (Select MAX(Finish_d) FROM K_A);
        return c;
END;
``` 

Запустить такую функцию не так просто

`Select type_t c from dual`

Если сделать так будет таблица, где единственный атрибут `[unsupported data type]`

Поэтому функция запускается так:

```
Select x.c.name_A, x.c.name_K, x.c.date_F
from
(Select type_t c from dual) x
```

Но... это не работает, пишет ошибку `ORA-06530: Reference to uninitialized composite`

Он говорит, что переменная не инициализирована.

Допустим, это можно сделать таким образом вернёмся и поправим в функции 4 строку:

`c F_TYPE:= F_TYPE('Hello','AH',sysdate);`

Запускаем функцию и... Всё работает. Только это не особо удобно делать, каждый раз объявлять какие-то значения?

Есть другой способ. Сделать конструктор в TYPE

Нужно поменять код TYPE на следующий:

```
CREATE OR REPLACE EDITIONABLE TYPE  "F_TYPE" AS OBJECT
( 
 name_A     VARCHAR2(255),
 name_K      VARCHAR2(255),
 date_F      DATE,
CONSTRUCTOR FUNCTION F_TYPE RETURN SELF AS RESULT
);
/
CREATE OR REPLACE EDITIONABLE TYPE BODY  "F_TYPE" 
AS
   CONSTRUCTOR FUNCTION F_TYPE RETURN SELF AS RESULT
   AS
   BEGIN
      RETURN;
   END;
END;
```

Создаём атрибуты и метод.

Теперь можно в функции изменить 4 строку на: 

`c F_TYPE:= F_TYPE();`

После этого запускаем тип - всё выполняется, как раньше.

Здесь мы явно объявили конструктор создания типа.