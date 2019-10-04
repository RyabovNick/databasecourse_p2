const Database = require('../config/database')
const { STRING, INTEGER, BLOB } = require('sequelize')
const types = require('./types')

const classes = Database.define(
  'classes',
  {
    name: STRING,
    description: BLOB('medium'),
    type_id: INTEGER
  },
  {
    underscored: true
  }
)

classes.belongsTo(types)

module.exports = classes
