require('dotenv').config()
const express = require('express')
const pool = require('./config/db')
const app = express()

app.route('/now').get(async (req, res) => {
  const pgclient = await pool.connect()
  const { rows } = await pgclient.query('SELECT now() as now')
  await pgclient.release()
  res.send(rows[0].now)
})

app.route('/user_order/:id').get(async (req, res) => {
  let pgclient
  try {
    // значение из URL
    pgclient = await pool.connect()
    const { id } = req.params
    const { rows } = await pgclient.query(`
      SELECT id, client_id, created_at
      FROM order_
      WHERE client_id = $1
      ORDER BY created_at DESC
    `, [id])
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

app.route('/make_order/:id').post(async (req, res) => {
  // TODO: получать id не из параметра, а из токена
  let pgclient
  try {
    pgclient = await pool.connect()
    const { id } = req.params
    const { rows } = await pgclient.query(`
    INSERT INTO order_ (client_id) VALUES ($1) RETURNING id
    `, [id])
    const orderID = rows[0].id
    res.send({
      order_id: orderID
    })

    // TODO: 
    // 1. Определиться со структурой, которую будем передавать
    // Возможно такая структура:
    // [
    //   {
    //     menu_id: 1,
    //     count: 2
    //   }
    // ]
    // 2. Всё выполнять в тразакции
    // 3. Определиться как считать стоимость заказа
    // 4. Добавить все продукты из заказа в order_menu

  } catch (err) {
    res.status(500).send({
      error: err.message
    })
    console.error(err)
  } finally {
    await pgclient.release()
  } 
})

app.listen(8080, () => {
  console.log('Server started on http://localhost:8080')
})