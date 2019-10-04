'use strict'
module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable('animals', {
      id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: Sequelize.INTEGER
      },
      name: {
        type: Sequelize.STRING
      },
      species_id: {
        type: Sequelize.INTEGER,
        references: {
          model: 'species',
          key: 'id'
        }
      },
      birth: {
        type: Sequelize.DATE
      },
      death: {
        type: Sequelize.DATE
      },
      weight: {
        type: Sequelize.DECIMAL(8, 3)
      },
      length: {
        type: Sequelize.DECIMAL(7, 2)
      },
      height: {
        type: Sequelize.DECIMAL(7, 2)
      },
      sex: {
        type: Sequelize.BOOLEAN
      },
      created_at: {
        allowNull: false,
        type: Sequelize.DATE
      },
      updated_at: {
        allowNull: false,
        type: Sequelize.DATE
      }
    })
  },
  down: (queryInterface, Sequelize) => {
    return queryInterface.dropTable('animals')
  }
}
