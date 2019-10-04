const Sequelize = require('sequelize')

const { PG_DATABASE, PG_USER, PG_PASSWORD, PG_HOST, PG_PORT } = process.env
const sequelize = new Sequelize(PG_DATABASE, PG_USER, PG_PASSWORD, {
  host: PG_HOST,
  port: PG_PORT,
  dialect: 'postgres',
  pool: {
    max: 20,
    idle: 30000,
    acquire: 60000
  }
})

module.exports = sequelize
