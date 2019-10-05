require('dotenv').config()
const Database = require('./config/database')
const types = require('./models/types')

const getClasses = async () => {
  const result = await types.findAll({
    attributes: ['name'],
  })

  console.log('result: ', result)
}

Database.authenticate().then(async () => {
  await Database.sync()
  const result = await getClasses()
  console.log('result: ', result)
})
