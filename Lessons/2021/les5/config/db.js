// Подключение к базе данных
const { Pool } = require('pg')
const pool = new Pool(
  {
    user: process.env.DB_USER,
    host: process.env.DB_HOST,
    database: process.env.DB_DATABASE,
    password: process.env.DB_PASSWORD,
    port: process.env.DB_PORT,
    max: 1,
  }
)

// экспортируем объект pool, чтобы из других файлов
// можно было использовать это подключение
module.exports = pool
