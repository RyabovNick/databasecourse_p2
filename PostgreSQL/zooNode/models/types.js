const Database = require('../config/database')
const { STRING } = require('sequelize')

const types = Database.define(
  'types',
  {
    name: STRING,
    description: BLOB('medium')
  },
  {
    underscored: true
  }
)

module.exports = types
