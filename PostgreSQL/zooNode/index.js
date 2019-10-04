require('dotenv').config()
const Database = require('./config/database')
const classes = require('./models/classes')

const getClasses = async () => {
  const result = await classes.findAll()

  console.log('result: ', result)
}

Database.authenticate().then(async () => {
  await Database.sync()
  const result = await getClasses()
  console.log('result: ', result)
})
