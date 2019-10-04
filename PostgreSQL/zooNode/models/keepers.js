const Database = require('../config/database')
const { STRING, DATE, SMALLINT } = require('sequelize')

const keepers = Database.define(
  'keepers',
  {
    name: STRING,
    surname: STRING,
    patronymic: STRING,
    birth: DATE,
    experience: SMALLINT
  },
  {
    underscored: true
  }
)

module.exports = keepers
