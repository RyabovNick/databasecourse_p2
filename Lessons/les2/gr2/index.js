// load .env to procces.env
require('dotenv').config()
const { Client } = require('pg')
const client = new Client({
  host: process.env.DB_HOST,
  port: process.env.DB_PORT,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_DATABASE,
})

// const name = 'name1'

client.connect()

// 1. Добавить новый автомобиль
// 2. Получить одного свободного менеджера
// 3. Назначить ему для продажи новый автомобиль
async function addCar() {
  const car = {
    brandID: 1,
    model: '420i',
    cost: 2500000,
    year: 2019,
    isAvailable: true,
  }

  try {
    // start transaction
    await client.query('BEGIN')
    const resCarID = await client.query(
      `
    INSERT INTO car (brand_id, model, cost, year_of_creation, is_available) VALUES
  ($1, $2, $3, $4, $5) RETURNING id
  `,
      [car.brandID, car.model, car.cost, car.year, car.isAvailable]
    )

    // throw 'ERROR!!!!!'

    const carID = resCarID.rows[0].id

    const resManagerID = await client.query(`
  SELECT * 
  FROM manager
  WHERE car_id IS NULL
  LIMIT 1
  `)
    const managerID = resManagerID.rows[0].id

    await client.query(
      `UPDATE manager
        SET car_id = $1
        WHERE id = $2`,
      [carID, managerID]
    )
    await client.query('COMMIT')
  } catch (err) {
    client.query('ROLLBACK')
  } finally {
    client.end()
  }
}

addCar()

// client
//     .query(
//         `
//   SELECT *
//   FROM test
//   WHERE name = $1
//   `,
//         [name]
//     )
//     .then((result) => console.log(result))
//     .catch((e) => console.error(e.stack))
//     .then(() => client.end())

// 1. По указанному ID менять статус на false
// 2. По указанному ID снижать цену на 10%

// 1. Получить все ID машин, старше 2018
// 2. Снизить цену на полученные авто на 5%

// ctrl+k+c
// ctrl+k+u
// client.query(`
//     SELECT *
//     FROM test
//     WHERE name = $1
//     `, [name], function (err, res) {
//   console.log(err, res)
//   client.end()
// })
// console.log(1)
