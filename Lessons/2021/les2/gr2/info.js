const a = {
    first: 1,
    second: 2
}

console.log(a.first)

let firstVar = a.first

console.log(firstVar)

const { first: firstL, second } = a
console.log(firstL)
console.log(second)

// Переменные окружения
console.log(process.env)
