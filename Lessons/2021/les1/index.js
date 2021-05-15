require('dotenv').config()
const { Pool } = require('pg')

const pool = new Pool({
	user: process.env.DB_USER,
	host: process.env.DB_HOST,
	database: process.env.DB_DATABASE,
	password: process.env.DB_PASSWORD,
	port: process.env.DB_PORT,
})

// pool.query('SELECT * FROM menu', (err, res) => {
// 	console.log(err, res)
// 	pool.end()
// })

/**
 * Создание нового заказа
 * 1. Создаём новый заказ и получаем его ID
 * 2. Подсчитать цену (TODO)
 * 3. Каждый товар из заказа добавить в таблицу order_menu
 */
async function createOrder() {
	const order = {
		cliendID: 1,
		menu: {
			id: 1,
			count: 2,
		},
	}

	try {
		const orderID = await pool.query(
			`INSERT INTO order_ (client_i) VALUES ($1) RETURNING id`,
			[order.cliendID]
		)
	} catch (error) {
		throw error
	} finally {
		pool.end()
	}
}

createOrder()
	.then(() => {
		console.log('success')
	})
	.catch((err) => {
		console.error(err)
	})
