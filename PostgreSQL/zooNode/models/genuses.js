const Database = require('../config/database')
const { STRING, INTEGER } = require('sequelize')
const families = require('./families')

const genuses = Database.define(
  'genuses',
  {
    name: STRING,
    description: BLOB('medium'),
    family_id: INTEGER
  },
  {
    underscored: true
  }
)

genuses.belongsTo(families)

module.exports = genuses
