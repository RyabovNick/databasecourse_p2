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
		// начинаем транзакцию
		await pool.query('BEGIN')

		const resOrderID = await pool.query(
			`INSERT INTO order_ (client_id) VALUES ($1) RETURNING id`,
			[order.cliendID]
		)

		const orderID = resOrderID.rows[0].id
		console.log('new order: ', orderID)

		// TODO: array меню с ценой
		const resMenu = await pool.query(
			`
		SELECT *
		FROM menu
		WHERE id = $1;`,
			[order.menu.id]
		)

		// throw new Error('ERRR!!!!!!')

		console.log('menu: ', resMenu.rows)

		const price = resMenu.rows[0].price * order.menu.count

		await pool.query(
			`INSERT INTO order_menu (order_id, menu_id, count, price) VALUES
			($1, $2, $3, $4);`,
			[orderID, order.menu.id, order.menu.count, price]
		)

		await pool.query('COMMIT')
	} catch (error) {
		await pool.query('ROLLBACK')
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
