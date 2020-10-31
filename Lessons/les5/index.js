require('dotenv').config()
const express = require('express')
const pool = require('./config/db')
// body parser, чтобы была возможность парсить body
const bodyParser = require('body-parser')
const bcrypt = require('bcryptjs')
const jwt = require('jsonwebtoken')

const secret = 'jwt_secret_value'

const app = express()
// чтобы парсить application/json
app.use(bodyParser.json())

// TODO API:
// 1) POST /sign_in - Войти в систему. 
//    Передаваться будет email, password.
//    Нужно проверить, что пользователь в таким email и password
//    существует.
// 2) GET /menu Получить меню. Без параметров 
//      (TODO: добавить пагинацию, сортировку и фильтры (поиск по цене, по весу))
// 3) DELETE /user_order/:id - (id - id заказа)

app.route('/now').get(async (req, res) => {
  const pgclient = await pool.connect()
  const { rows } = await pgclient.query('SELECT now() as now')
  await pgclient.release()
  res.send(rows[0].now)
})

/**
 * checkAuth валидирует токен, 
 * в случае успеха возвращает payload
 * @param {*} req 
 */
async function checkAuth(req) {
  const authHeader = req.headers.authorization

  let token 
  if (authHeader) {
    const h = authHeader.split(' ')
    if (h[0] !== 'Bearer') {
      throw new Error('Allowed only Bearer token')
    }

    token = h[1]
  } else {
    throw new Error('Token not found')
  }

  try {
    return jwt.verify(token, secret)
  } catch (err) {
    throw err
  }
}

// Все заказы конкретного пользователя
// id пользователя берётся из токена
app.route('/user_order').get(async (req, res) => {
  let tokenPaylod
  try {
    tokenPaylod = await checkAuth(req)
  } catch (err) {
    res.status(401).send({
      error: err.message
    })
    return
  }

  let pgclient
  try {
    // значение из URL
    pgclient = await pool.connect()
    const { rows } = await pgclient.query(`
      SELECT id, client_id, created_at
      FROM order_
      WHERE client_id = $1
      ORDER BY created_at DESC
    `, [tokenPaylod.id])
    res.send(rows)
  } catch (err) {
    res.status(500).send({
      error: err.message
    })
    console.error(err)
  } finally {
    // Не забываем всегда закрывать соединение с базой
    await pgclient.release()

  }
})

// Сделать новый заказ
// Структура body:
// [
//   {
//     menu_id: 1,
//     count: 2
//   }
// ]
app.route('/make_order/:id').post(async (req, res) => {
  // TODO: получать id не из параметра, а из токена

  // TODO: обработать ошибку, когда подключиться не удалось
  let pgclient = await pool.connect()
  try {
    const { id } = req.params

    // открываем транзакцию
    await pgclient.query('BEGIN')

    // Создали заказ и получили его ID
    const { rows } = await pgclient.query(`
    INSERT INTO order_ (client_id) VALUES ($1) RETURNING id
    `, [id])
    const orderID = rows[0].id

    // делаем цикл по body
    // чтобы подготовить запрос на получение цены
    // по каждому товару из заказа

    // параметры для подготовки IN запроса
    // пример: IN ($1,$2,$3)
    let params = [] // ["$1", "$2", "$3"]
    let values = [] // [1, 2, 3]
    for (const [i, item] of req.body.entries()) {
      params.push(`$${i+1}`)
      values.push(item.menu_id)
    }

    // Получить стоимость из меню
    const { rows: costQueryRes } = await pgclient.query(`
      SELECT id, price::numeric
      FROM menu
      WHERE id IN (${params.join(',')})
    `, values)

    // мы хотим содать новую переменную, которая
    // будет включать тоже самое, что и
    // входной body, только с вычисленной ценой
    let orderWithCost = []
    // для этого надо пройтись по каждому элементу
    // в body
    for (const item of req.body) {
      // и для каждого элемента найти цену в costQuery
      // полученном при помощи запроса
      let cost = null
      for (const i of costQueryRes) {
        // ищем совпадение id в costQuery
        // с menu_id переданном в body
        if (i.id === item.menu_id) {
          cost = i.price
        }
      }

      // тут cost либо null, либо с значением цены
      // и если cost null, означает, что такого товара
      // в таблице menu не найдено, т.е. ошибка
      // Нам надо сделать rollback, вернуть сообщение клиенту
      if (!cost) {
        throw new Error(`Not found in menu: ${item.menu_id}`)
      }

      orderWithCost.push({
        ...item,
        cost: cost * item.count // найденную стоимость на кол-во
      })
    }

    // добавляем все продукты заказа в order_menu
    // оптимальный вариант, это сгенерировать один
    // INSERT, который сразу добавит всё в таблицу
    // order_menu (как мы делали раньше)
    // Но тут попробуем сделать с Promise.all
    // т.е. отправить одновременно в базу все запросы
    // а уже после отправки ждать выполнение их всех
    // вместе.
    let promises = []
    for (const item of orderWithCost) {
      promises.push(pgclient.query(
        `INSERT INTO order_menu (order_id, menu_id, count, price) 
          VALUES ($1, $2, $3, $4);`,
        [orderID, item.menu_id, item.count, item.cost]
      ))
    }

    // ждём, пока выполнятся все запросы
    await Promise.all(promises)

    // коммитим изменения в базе
    await pgclient.query('COMMIT')
    res.send({
      order_id: orderID
    })
  } catch (err) {
    // Всегда, если мы попадаем в catch, то
    // откатываем транзакцию
    await pgclient.query('ROLLBACK') 
    // и отправляем клиенту, что произошла ошибка
    res.status(500).send({
      error: err.message
    })
    console.error(err)
  } finally {
    // освобождаем соединение с postgresql
    await pgclient.release()
  } 
})

app.route('/sign_in').post(async (req, res) => {
  const {
    email,
    password
  } = req.body

  try {
    const { rows } = await pool.query(`
    SELECT id, email, password
    FROM client
    WHERE email = $1
    `, [email])

    // если пользователь с таким email
    // не найден
    if (rows.length == 0) {
      res.status(401).send({
        error: 'User not found'
      })
      return
    }

    // проверяем правильность пароля
    const isValid = await bcrypt.compare(password, rows[0].password)
    if (!isValid) {
      res.status(401).send({
        error: 'Invalid password'
      })
      return
    }

    // если правильность введённых данных пользователем
    // подтверждена
    const token = jwt.sign({
      id: rows[0].id,
      email: rows[0].email
    }, secret, {
      expiresIn: "1d",
    })

    res.send({
      token
    })

  } catch (err) {
    res.status(500).send({
      error: err.message
    })
  }
})

// Зарегистрироваться
app.route('/sign_up').post(async (req, res) => {
  // Если какой-то из параметров не будет передан, то
  // будет SQL ошибка (NOT NULL contraint)
  // По хорошему, нам надо тут проверить, что 
  // параметры, которые не могут быть NULL переданы
  const { 
    name,
    address,
    phone,
    email, 
    password 
  } = req.body

  let pgclient = await pool.connect()
  try {
    const hash = await bcrypt.hash(password, 8)

    const { rows } = await pgclient.query(`
    INSERT INTO client (name, address, phone, email, password)
    VALUES ($1,$2,$3,$4,$5) RETURNING id;
    `, [name, address, phone, email, hash])

    // TODO:
    // 2) Добавить JWT и генерить токен, возвращать в ответе на запрос
    // вместе с id. В токен в payload добавить id

    res.send({
      id: rows[0].id
    })
  } catch (err) {
    res.status(500).send({
      error: err.message
    })
    console.error(err)
  } finally {
    // освобождаем соединение с postgresql
    await pgclient.release()
  }
})

app.listen(8080, () => {
  console.log('Server started on http://localhost:8080')
})