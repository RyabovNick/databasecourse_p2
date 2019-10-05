const Database = require('../config/database')
const { STRING, INTEGER } = require('sequelize')
const species = require('./species')

const animals = Database.define(
  'animals',
  {
    name: STRING,
    species_id: INTEGER,
    birth: DATE,
    death: DATE,
    weight: DECIMAL,
    length: DECIMAL,
    height: DECIMAL,
    sex: BOOLEAN,
  },
  {
    underscored: true,
  },
)

animals.belongsTo(species)

module.exports = animals
