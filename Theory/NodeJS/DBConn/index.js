const express = require("express");
const app = express();
//подключаем ранее созданный файл
const pool = require("./config");

//по заданному адресу будет выполнять подключение к БД
//и возврат клиенту всех данных из таблицы
app.route("/get").get((request, response) => {
  //берём свободное подключение
  pool.getConnection((err, connection) => {
    //если возникнет ошибка - будет выведена в консоль
    if (err) throw err;
    //делаем запрос, параметры ошибка и результат
    connection.query("Select * from tablename", (error, result) => {
      //если возникнет ошибка - будет выведена в консоль
      if (error) throw error;
      //отправить клиенту результат выполнения запроса
      response.send(result);
    });
    //Обязательно! Необходимо освободить соединение, иначе через 100 штук приложение упадёт
    connection.release();
  });
});

app.listen(8080, () => {
  console.log("Server started!");
});
