# Search engine

## Elasticsearch

Использование elasticsearch и kibana (GUI)

Мы будем добавлять некоторые документы,
искать и удалять их. После этого будем использовать
датасет Шекспира для более глубокого поиска и
агрегации.

Мы не будем рассматривать то, что связано с конфигурацией,
лучшими практиками. Полученную информацию можно использовать
в качестве представления того, какие возможности предоставляет
Elasticsearch и как он может соответствовать потребностям.

Взаимодействие с Elasticsearch осуществляется при
помощи REST API. Для удобства взаимодействия и
визуализации данных мы ещё используем kibana -
графический пользовательский интерфейс

REST API:  
позволяет отправлять разные виды запросов  
GET - получить
POST - добавить
PUT - изменить
DELETE - удалить

Допустим, запрос

`GET localhost:9200`

Вернёт нечто подобное:

```
{
    "name": "myNode",
    "cluster_name": "myCluster",
    "cluster_uuid": "ysWbrK0fScqfRx-HBSRDyA",
    "version": {
        "number": "6.3.1",
        "build_flavor": "default",
        "build_type": "deb",
        "build_hash": "eb782d0",
        "build_date": "2018-06-29T21:59:26.107521Z",
        "build_snapshot": false,
        "lucene_version": "7.3.1",
        "minimum_wire_compatibility_version": "5.6.0",
        "minimum_index_compatibility_version": "5.0.0"
    },
    "tagline": "You Know, for Search"
}
```

Для работы с API можно использовать удобные продукты,
например POSTMAN.

```
POST localhost:9200/accounts/person/1 
{
    "name" : "John",
    "lastname" : "Doe",
    "job_description" : "Systems administrator and Linux specialit"
}
```

Вышеуказанный запрос создаст person с id 1
и вернёт ответ -

```
{
    "_index": "accounts",
    "_type": "person",
    "_id": "1",
    "_version": 1,
    "result": "created",
    "_shards": {
        "total": 2,
        "successful": 1,
        "failed": 0
    },
    "created": true
}
```

Теперь если мы обратимся к

`GET localhost:9200/accounts/person/1`

Получим такой ответ

```
{
    "_index": "accounts",
    "_type": "person",
    "_id": "1",
    "_version": 1,
    "found": true,
    "_source": {
        "name": "John",
        "lastname": "Doe",
        "job_description": "Systems administrator and Linux specialit"
    }
}
```

Мы допустили ошибку - specialit

Давайте исправим её.
Это можно сделать при помощи _update

```
POST localhost:9200/accounts/person/1/_update
{
      "doc":{
       "job_description" : "Systems administrator and Linux specialist"
     }
}
```

Теперь посмотрим, что у нас лежит

```
{
    "_index": "accounts",
    "_type": "person",
    "_id": "1",
    "_version": 2,
    "found": true,
    "_source": {
        "name": "John",
        "lastname": "Doe",
        "job_description": "Systems administrator and Linux specialist"
    }
}
```

Ошибка исправлена!

Давайте добавим ещё одну персону:

```
POST localhost:9200/accounts/person/2
{
    "name" : "John",
    "lastname" : "Smith",
    "job_description" : "Systems administrator"
}
```

Главное назначение данной БД - поиск.

Давайте сделаем один. Для этого есть
определённый формат /_search?q=something  
где something - то, что мы хотим найти

```
GET localhost:9200/_search?q=john
```

Поиск вернёт всех найденных Джонов

```
{
    "took": 58,
    "timed_out": false,
    "_shards": {
        "total": 5,
        "successful": 5,
        "failed": 0
    },
    "hits": {
        "total": 2,
        "max_score": 0.2876821,
        "hits": [
            {
                "_index": "accounts",
                "_type": "person",
                "_id": "2",
                "_score": 0.2876821,
                "_source": {
                    "name": "John",
                    "lastname": "Smith",
                    "job_description": "Systems administrator"
                }
            },
            {
                "_index": "accounts",
                "_type": "person",
                "_id": "1",
                "_score": 0.28582606,
                "_source": {
                    "name": "John",
                    "lastname": "Doe",
                    "job_description": "Systems administrator and Linux specialist"
                }
            }
        ]
    }
}
```

Можно использовать различные поиски:

`GET localhost:9200/_search?q=smith`

`GET localhost:9200/_search?q=job_description:john`

`GET localhost:9200/accounts/person/_search?q=job_description:linux`

Мы ещё не делали удалёние. Выполняется оно так:

`DELETE localhost:9200/accounts/person/1`

`DELETE localhost:9200/account`

В итоге, в этом кратком обзоре мы:

1. Добавили документ и был создан индекс
2. Получили документ
3. Изменили документ
4. Добавили второй документ
5. Выполнили различный поиск
6. Удалили один документ
7. Удалили целый индекс

Добавим датасет  
<https://www.elastic.co/guide/en/kibana/6.3/tutorial-load-dataset.html>

`GET localhost:9200/shakespeare/_search`

Вернёт 10 записей

```
POST localhost:9200/shakespeare/scene/_search/
{
    "query":{
     "bool": {
         "must" : [
             {
                 "match" : {
                     "play_name" : "Antony"
                 }
             },
             {
                 "match" : {
                     "speaker" : "Demetrius"
                 }
             }
         ]
     }
    }
}
```

Если мы хотим создать некую аналитку, то
нам понадобятся агрегации. Они позволяют проводить
более глубокое понимание данных.
Например, сколько разных пьес в наших данных,
сколько сцен в среднем на работу. Какие работы с
наибольшим количеством сцен.

Вернёмся к моменту, когда мы создали Shakespeare
Index. В Elastic мы может создавать индексы,
определяющие тип данных для разных полей. Они
могут быть: числовые, ключевые, текстовые и много
других типов

<https://www.elastic.co/guide/en/elasticsearch/reference/6.3/mapping-types.html>

Типы данных, которые могут иметь индексы определяются
через сопоставления (mappings)

<https://www.elastic.co/guide/en/elasticsearch/reference/6.3/mapping.html>

Mappings - процесс определения как документ и
его поля хранятся и индексируются.  
При помощи этого определяется:

* какие string поля должны обрабатываться как
полные текстовые поля
* какие поля содержат цифры, даты или геолокацию
* нужно ли индексировать значения всех полей
в документе в _all поле.
* формат даты
* настраиваемые правила для управления отображением
для динамически добавленных полей.

В данном случае мы не создавали никаких индексов,
поэтому Elastic решил какой тип каждого поля.

Тип text был выбран для текстовых полей. Этот тип
анализируется, вот что позволило найти play_name.

По умолчанию мы не можем делать агрегации в анализиуемых
полях.

Вопрос в том как собираемся делать агрегации, если
поля недействительны для их выполнения? Elastic,
когда определил тип каждого поля также добавил не
анализируемую версию текстовых полей (keyword) как
раз для случая, если мы захотим делать агрегации, сортировки,
скрипты. Мы просто может использовать `play_name.keyword`.

```
POST localhost:9200/shakespeare/_search
{
    "size":0,
    "aggs" : {
        "Total plays" : {
            "cardinality" : {
                "field" : "play_name.keyword"
            }
        }
    }
}
```
