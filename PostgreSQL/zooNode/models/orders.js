const Database = require('../config/database')
const { STRING, INTEGER, BLOB } = require('sequelize')
const classes = require('./classes')

const orders = Database.define(
  'orders',
  {
    name: STRING,
    description: BLOB('medium'),
    class_id: INTEGER,
  },
  {
    underscored: true,
  },
)

orders.belongsTo(classes)

module.exports = orders
