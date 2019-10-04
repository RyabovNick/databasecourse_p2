require('dotenv').config()

const { PG_DATABASE, PG_USER, PG_PASSWORD, PG_HOST, PG_PORT } = process.env
// названия берутся из NODE_ENV
module.exports = {
  development: {
    username: PG_USER,
    password: PG_PASSWORD,
    database: PG_DATABASE,
    host: PG_HOST,
    port: PG_PORT,
    dialect: 'postgres'
  }
}
