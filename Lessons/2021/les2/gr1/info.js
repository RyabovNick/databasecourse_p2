const obj = {
    first: 1,
    second: 5,
}

const first = obj.first
const secondA = obj.second
console.log(first)
console.log(secondA)

const { first: firstL, second } = obj
console.log(firstL)
console.log(second)

console.log(process.env)
