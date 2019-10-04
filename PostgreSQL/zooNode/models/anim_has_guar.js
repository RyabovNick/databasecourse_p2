const Database = require('../config/database')
const { INTEGER } = require('sequelize')
const animals = require('./animals')
const guardians = require('./guardians')

const anim_has_guar = Database.define(
  'anim_has_guar',
  {
    guardian_id: {
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
    }
  },
  {
    underscored: true
  }
)

anim_has_guar.belongsTo(animals)
anim_has_guar.belongsTo(guardians)

module.exports = anim_has_guar
