const http = require('http')

const server = http.createServer((req, res) => {
  console.log(req.url, req.method)

  // работаем только когда url /get
  if (req.url === '/get') {
    res.writeHead(200, { 'Content-Type': 'text/plain' })
    res.end('Its get')
    // важно не забыть тут return,
    // иначе код пойдёт дальше, т.е. res.end - отправит
    // ответ клиенту, но функция сама не завершилась
    // т.е. дальше наша функция попробует отправить снова
    // сообщение клиенту, но соединение уже завершено
    return
  }

  res.writeHead(200, { 'Content-Type': 'text/plain' })
  res.end('Hello world')
})

// server запущен на localhost:8080
server.listen(8080)