const array = [1, 2, 3, 4, 5, 2, 5, 6, 7]

for (const iterator of array) {
    console.log(iterator);
}

/**
 * Вычисляет сумму массива
 * @param {*} arr - текущий массив
 * @param {*} acc - сумма
 */
const sum = (arr, acc) => {
    if (arr.length === 0) return acc

    return sum(arr.slice(1), acc + arr[0])
}

let arraySum = 0
console.log(sum(array, arraySum))

const reduceSum = array.reduce((acc, currentValue) => acc + currentValue)
// acc  currentValue    return
// 0    1               1
// 1    2               3
// 3    3               9
// 9    4               13
const reduceSum1 = array.reduce(function (acc, currentValue) { return acc + currentValue })
console.log(reduceSum)