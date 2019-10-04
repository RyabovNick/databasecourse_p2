const Database = require('../config/database')
const { STRING } = require('sequelize')

const guardians = Database.define(
  'guardians',
  {
    name: STRING
  },
  {
    underscored: true
  }
)

module.exports = guardians
