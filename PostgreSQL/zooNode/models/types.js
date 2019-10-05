const Database = require('../config/database')
const { STRING, BLOB } = require('sequelize')

const types = Database.define(
  'types',
  {
    name: STRING,
    description: BLOB('medium'),
  },
  {
    underscored: true,
  },
)

module.exports = types
