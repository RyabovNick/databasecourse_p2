const Database = require('../config/database')
const { STRING, INTEGER, BLOB } = require('sequelize')
const orders = require('./orders')

const orders = Database.define(
  'orders',
  {
    name: STRING,
    description: BLOB('medium'),
    order_id: INTEGER,
  },
  {
    underscored: true,
  },
)

orders.belongsTo(families)

module.exports = orders
