const pool = require('../../config/db')

/**
 * Возвращаёт всё меню или выполняет поиск по названию
 *
 * TODO: а) пагинация б) сортировка в) фильтр по цене / весу
 * @param {string} name - название продукта
 */
async function findMenu(name) {
  // Используем WHERE 1=1, чтобы не
  // делать условие добавления WHERE для
  // каждого фильтра
  let query = `
  SELECT *
  FROM menu
  WHERE 1=1
  `
  const values = []

  if (name) {
    query += 'AND name ILIKE $1'
    values.push(`%${name}%`)
  }

  // TODO: решить проблему, с $1 <- параметром
  // в запросе. (Например счётчик)

  const { rows } = await pool.query(query, values)
  return rows
}

module.exports = {
  findMenu,
}
