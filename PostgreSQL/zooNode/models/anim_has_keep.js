const Database = require('../config/database')
const { INTEGER, DATE } = require('sequelize')
const animals = require('./animals')
const keepers = require('./keepers')

const anim_has_keep = Database.define(
  'anim_has_keep',
  {
    keepere_id: {
      allowNull: false,
      autoIncrement: false,
      primaryKey: true,
      type: INTEGER
    },
    animal_id: {
      allowNull: false,
      autoIncrement: false,
      primaryKey: true,
      type: INTEGER
    },
    start: DATE,
    finish: DATE
  },
  {
    underscored: true
  }
)

anim_has_keep.belongsTo(animals)
anim_has_keep.belongsTo(keepers)

module.exports = anim_has_keep
