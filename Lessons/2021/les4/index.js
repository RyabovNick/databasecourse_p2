const express = require('express')
const app = express()

// route сработает на http://localhost:8080/get
app.route('/get').get((req, res) => {
  res.send('Hello world')
})

// route с параметром :group
// сработает на http://localhost:8080/group/123 (любые значения вместо 123)
app.route('/group/:group').get((req, res) => {
  // из params объекта мы можем достать свойство
  const group = req.params.group
  res.send(`Hello group!: ${group}`)
})

app.listen(8080, () => {
  console.log('Server started on http://localhost:8080')
})