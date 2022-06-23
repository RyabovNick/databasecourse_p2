# Functions

1. Что такое функиця?
2. Чем функция отличается от процедуры?
3. Вызов функции
4. Возвращаемые параметры

**Функции** отличаются от процедур тем, что возвращают какое-то значение и вызываются только из других предложений SQL. Синтаксически это отличие отражено в спецификации функции. В остальном, функция создается по тем же правилам, что и процедура. Например, оператор создания функции, возвращающей максимальную цену среди всех товаров может выглядеть следующим образом:

```
CREATE FUNCTION get_max_price
RETURN NUMBER
AS
Max_price NUMBER;
BEGIN
SELECT max(price) INTO max_price FROM pricelist;
RETURN max_price;
END;
```

Заметьте, что в спецификации функции присутствует служебное слово RETURN, после которого указывается тип возвращаемого значения. Размер типа возвращаемого значения также как и размер типа входных и выходных параметров не должен быть конкретизирован.

В секции объявлений объявлена одна переменная max_price типа NUMBER, которая используется для хранения максимальной цены, которую возвращает запрос. Для возврата значения из функции в вызвавшее ее SQL предложение используется оператор RETURN.

Также стоит отметить некоторое отличие в синтаксисе оператора SELECT, а именно присутствие служебного слова INTO. Внутри процедур, функций и триггеров нельзя использовать оператор SELECT без служебного слова INTO. После INTO в запросе должна быть указана одна или несколько переменных, в которые и запишутся результаты выполнения запроса. Естественно, что эти переменные должны быть предварительно объявлены в секции объявлений.

**Вызов функций:**

Во время отладки функции для проверки ее работоспособности используется следующая конструкция:

```
SELECT имя_функции (входные_параметры) FROM DUAL;
```

**Удаление функций:**

Удаление функция производится посредством оператора _DROP FUNCTION_ имя_функции;