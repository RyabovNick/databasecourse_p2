const pool = require('../../config/db')

/**
 * Возвращаёт всё меню или выполняет поиск по названию
 *
 * TODO: а) пагинация б) сортировка в) фильтр по цене / весу
 * @param {string} name - название продукта
 * @param {number} price - максимальная цена
 */
async function findMenu(name, price) {
  // Используем WHERE 1=1, чтобы не
  // делать условие добавления WHERE для
  // каждого фильтра
  let query = `
  SELECT *
  FROM menu
  WHERE 1=1
  `
  const values = []

  let counter = 1
  if (name) {
    query += `AND name ILIKE $${counter}`
    values.push(`%${name}%`)
    counter++
  }

  if (price) {
    query += ` AND price < $${counter}`
    values.push(price)
    counter++
  }

  const { rows } = await pool.query(query, values)
  return rows
}

module.exports = {
  findMenu,
}
