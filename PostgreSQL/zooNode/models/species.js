const Database = require('../config/database')
const { STRING, INTEGER } = require('sequelize')
const genuses = require('./genuses')

const species = Database.define(
  'species',
  {
    name: STRING,
    description: BLOB('medium'),
    genus_id: INTEGER,
  },
  {
    underscored: true,
  },
)

species.belongsTo(genuses)

module.exports = species
