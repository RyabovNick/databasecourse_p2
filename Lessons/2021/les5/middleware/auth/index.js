const jwt = require('jsonwebtoken')
const secret = 'jwt_secret_value'

/**
 * checkAuth валидирует токен,
 * в случае успеха добавляет в req client
 * @param {*} req
 */
function checkAuth(req, res, next) {
  const authHeader = req.headers.authorization

  let token
  if (authHeader) {
    const h = authHeader.split(' ')
    if (h[0] !== 'Bearer') {
      res.status(401).send({
        error: 'Allowed only Bearer token',
      })
      return
    }

    token = h[1]
  } else {
    res.status(401).send({
      error: 'Token not found',
    })
    return
  }

  // чтобы следующие после middleware функции
  // могли использовать переменные, полученные тут
  // добавим client в req
  try {
    req.client = jwt.verify(token, secret)
  } catch (err) {
    res.status(401).send({
      error: err.message,
    })
    return
  }

  next()
}

module.exports = checkAuth
