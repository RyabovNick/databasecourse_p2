'use strict';
module.exports = (sequelize, DataTypes) => {
  const animals = sequelize.define('animals', {
    name: DataTypes.STRING,
    species_id: DataTypes.INTEGER,
    birth: DataTypes.DATE,
    death: DataTypes.DATE,
    weight: DataTypes.DECIMAL,
    length: DataTypes.DECIMAL,
    height: DataTypes.DECIMAL,
    sex: DataTypes.BOOLEAN
  }, {
    underscored: true,
  });
  animals.associate = function(models) {
    // associations can be defined here
  };
  return animals;
};