const arr = [1,2]
var varv
console.log(arr)

arr.push(3)
console.log(arr)

console.log(varv)
var varv = 1

const obj = {
    a: 1,
    b: 2,
    c: null,
    d: {
        dd: 1,
        cc: 2
    }
}

console.log(obj.d.dd)

// console.log(obj.f.d)

const fl = ""
if (fl) {
    console.log("fl exists")
} 

if (obj) {
    console.log("obj exists")
}

// const t = obj.f.d
// console.log(t)

const a = obj.a != 1 ? 3 : 5 

let b
if (obj.a != 1) {
    b = 3
} else {
    b = 5
}


let one = 1
let str = '1'
let two = 2

console.log(one + two)

one = true

console.log(one + two)

one = '1'

console.log(one + two)

two = true

console.log(one + two)

const student = {
    name: 'Иван',
    surname: 'Иванов'
}

// Имя: Иван, Фамилия: Иванов

console.log('Имя: ' + student.name + ', Фамилия: ' + student.surname)
console.log(`Имя: ${student.name}, Фамилия: ${student.surname}`)

const t1 = {
    a: 1
}

const t2 = {
    a: '1'
}

let v = {
    a: '1'
}

if (t1.a === t2.a) {
    console.log('Its equal')
} else {
    console.log('It isnt equal')
}

const arr1 = [1,2,3,4,56,7]
const obj1 = {
    a: 'test',
    b: 'test1'
}

const objArray = [
    {
        a: 1,
        b: 2
    }, 
    {
        a: 4,
        b: 5
    }
]

for (const [i, v] of arr1.entries()) {
    console.log(i)
    console.log(v)
}

arr1.forEach(element => {
    console.log(element)
});

for (const key in obj1) {
    console.log(key)
    console.log(obj1[key])
}

const arr2 = [1,4,6,3,4,5,9,2,1]

// arr3 = [1,3,2,1]
let arr3 = []
for (const v of arr2) {
    // if (v <= 3) arr3.push(v)
    if (v <= 3) arr3 = [...arr3, v]
}
console.log(arr3)

let arr4 = arr2.filter(item => item <= 3)
console.log(arr4)