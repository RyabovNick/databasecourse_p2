# Создание API при помощи встроенных средств apex.oracle.com

Oracle предоставляет возможности легкого создания API. Подобное решение можно использовать в учебных целях.

О том, что такое API можно почитать тут (до части Tutorial, в нем в кратце описано как сделать API собственными средствами) [API](https://bitbucket.org/For_Victory/unidubnadb/src/master/API/API.md)

## Инструкция

SQL Workshop -> RESTful Services

ORDS RESTful Services

Далее необходимо нажать на кнопку "Register Schema with ORDS" и оставить можно все настройки по умолчанию.

После этого создастся один стандартный модуль с employees. Можно посмотреть его.

Нажмите на кнопку Modules и создайте новый. Придумайте название и можно поменять другие параметры.

Дальше находясь в нужном модуле создайте Template, указать название, если есть входные параметры их можно указать в таком виде animals/:id в таком случае :id можно использовать внутри запроса

```
Select *
from animals
where id = :id
```

В таком случае при обращении по адресу animals/1 будет получена вся информация о животном с заданным id (т.е. 1)

Для созданного Template необходимо создать Handler. В нём необходимо указать метод GET - получение. PUT - изменение. POST - добавление. DELETE - удаление.

Для 1-го примера сделаем GET.

Также необходимо выбрать Source Type. Я не буду сильно останавливаться на отличии каждого из вариантов - более подробно о них можно почитать, нажав на кнопку с "?".

Для нас подходит вариант Collection Query, который вернёт всё в формате JSON. Ну и PL/SQL, может пригодится - запускает только анонимные блоки.

Напишите любой запрос, допустим, получения всех животных (если делали без входного параметра Template), либо для конкретного животного (если Template с входным параметром).

После создания GET запроса скопируйте и введите в адресную строку или используя дополнение Rest Client для Visual Studio Code или используя ПО Postman или любой другой вариант.

Посмотрите на результат - в формате JSON.

Как это использовать можно посмотреть в папке 1. Hello World

Далее можно создать POST, PUT, DELETE API. Их отличие в том, что их можно реализовать только при помощи PL/SQL.

Например, создайте к созданному Template animals Handler POST.

### PUT

Для создания PUT запроса есть единственный возможный Source Type - PL/SQL.

Для описания инструкций используется анонимный блок, т.е.

```
begin
...
end;
```

Если необходимо, то зона Declare. Можно использовать ранее написанные функции, процедуры, пакеты.

Важная особенность этой части - использование параметров:

1. name: X-APEX-FORWARD (стандартное название переменной), bind variable: location (можете назвать её как-нибудь по-другому), source type: Http Header, Access method: OUT, Data Type: String. Этот параметр возвращает результат запроса клиенту (в JSON формате).
2. name: X-APEX-STATUS-CODE, bind variable: status, source type: Http Header, Access method: OUT, Data Type: Integer. Этот параметр возвращает статус запроса (обычный HTTP код, можно найти. Успешно - 200)
3. name: ID, bind variable: id, source type: Http Header, method: IN, Data Type: Integer. Этот параметр - входной.

Готовый код в качестве примера

```
begin
    update keepers
    set
        name = nvl(:name, name),
        surname = nvl(:surname, surname),
        patronymic = nvl(:patronymic, patronymic),
        birth = nvl(:birth, birth),
        experience = nvl(:experience, experience)
    where
        id = :id;
    :status := 200;
    :location := :id;
end;
```

Функция nvl - используется для того, чтобы в случае неизменения какого-то атрибута использовалось старое значение. Это позволяет не писать кучу update для каждого атрибута.

`:status:= 200` Если запрос выполнился хорошо, то вернуть в ответе status - 200.

`:location := :id` насколько я понял, тут используется ранее написанный GET запрос для keepers/:id - получить по заданному id всё о конкретном работнике. Если GET запрос не написан - возникнут ошибки. Ответ на запрос записывается в переменную :location и возвращает результат в примерно таком виде:

```
HTTP/1.1 200 OK
Date: Tue, 30 Oct 2018 09:32:29 GMT
Server: Oracle-HTTP-Server-11g
X-ORACLE-DMS-ECID: 005UMvTGiba9pYKqESi8US0000nI0000wq
ETag: "KUQkMSp6Tif4qJrwYLRhURAjBXleFaKBA1/3RwjmQx+/g1C2IM6EpCZMnfnrl5uu+jVsNj53JLrK93waeZcFZg=="
Content-Location: https://apex.oracle.com/pls/apex/for_victory1/zoo/keepers/67
X-ORACLE-DMS-RID: 0:1
Vary: Accept-Encoding
Keep-Alive: timeout=5, max=100
Connection: Keep-Alive
Transfer-Encoding: chunked
Content-Type: application/json
Content-Language: en

{
  "items": [{
    "id": 67,
    "name": "What",
    "surname": "Worker2",
    "patronymic": "valentine1",
    "birth": null,
    "experience": 7
  }],
  "hasMore": false,
  "limit": 25,
  "offset": 0,
  "count": 1,
  "links": [{
    "rel": "self",
    "href": "https://apex.oracle.com/pls/apex/for_victory1/zoo/keepers/67"
  }, {
    "rel": "edit",
    "href": "https://apex.oracle.com/pls/apex/for_victory1/zoo/keepers/67"
  }, {
    "rel": "describedby",
    "href": "https://apex.oracle.com/pls/apex/for_victory1/metadata-catalog/zoo/keepers/item"
  }, {
    "rel": "first",
    "href": "https://apex.oracle.com/pls/apex/for_victory1/zoo/keepers/67"
  }]
}
```

Для отправки запросов PUT, POST можно использовать различный софт. Например, расширение для Visual Studio Code - REST Client (примеры запуска находятся в 0. Insert animals). Также можно использовать программу Postman, cURL - по вашему желанию. На парах основное средство - Rest client for VS code.

### POST

Делается для /keepers (т.е. без входного параметра).

Можно использовать для передачи статуса запроса тот же параметр X-APEX-STATUS-CODE (как в PUT).

Также, если хотите, чтобы запрос вернул что-то (например ID добавленной записи (самый распространённый случай), то можно сделать OUT параметр, например ID и source type - URI.

В самом коде необходимо выполнить добавление новой строки, вернуть ID этой строки.

```
DECLARE
   id keepers.id%type;
BEGIN
  INSERT INTO keepers (name, surname)
     VALUES (:name, :surname)
  RETURNING ID INTO id;
  :id := id;
  :status := 201;
END;
```

Обратите внимание на дополнительную строчку в INSERT - `RETURNING ID INTO id;`, которая возвращает заданный атрибут (в данном случае ID и кладёт его в переменную id.

`:id := id;` :id - это параметр, выходной, в который положится то, что содержит переменная id.

```
HTTP/1.1 201 Created
Date: Tue, 30 Oct 2018 10:02:06 GMT
Server: Oracle-HTTP-Server-11g
X-ORACLE-DMS-ECID: 005UMx7E4CJ9pYKqESi8US0000nI00008g
X-ORACLE-DMS-RID: 0:1
Vary: Accept-Encoding
Keep-Alive: timeout=5, max=100
Connection: Keep-Alive
Transfer-Encoding: chunked
Content-Type: application/json
Content-Language: en

{
  "ID": "123"
}
```

После выполнения запроса результат будет примерно такой. Пример отправки запроса в файле post_keepers.http

# Delete

Этот метод создаём в keepers/:id, где вместо :id подставляем номер, который будет удалён.

Используем параметр X-APEX-STATUS-CODE для отправки статуса.

```
begin
    Delete from keepers
    where id = :id;
    :status := 200;
end;
```

Просто удаляем строку с заданным id.

Ответ на запрос (вариант запуска можно найти в delete_keepers.http) будет выглядеть примерно так:

```
HTTP/1.1 200 OK
Date: Tue, 30 Oct 2018 10:27:00 GMT
Server: Oracle-HTTP-Server-11g
X-ORACLE-DMS-ECID: 005UMyWFIF69pYKqESi8US0000nI0001mN
X-ORACLE-DMS-RID: 0:1
Content-Length: 0
Keep-Alive: timeout=5, max=100
Connection: Keep-Alive
Content-Type: text/plain
Content-Language: en
```
