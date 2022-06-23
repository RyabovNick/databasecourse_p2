/**Синхронный стиль */
// прямой стиль
function add(a, b) {
  return a + b;
}

// w/ callback
function add(a, b, callback) {
  callback(a + b);
}
console.log("not async");
console.log("before");
add(1, 2, result => console.log("Result: " + result));
console.log("after");
console.log("_________");

/**Асинхронный стиль */
function additionAsync(a, b, callback) {
  setTimeout(() => callback(a + b), 100);
}
console.log("async");
console.log("before");
additionAsync(1, 2, result => console.log("Result: " + result));
console.log("after");
console.log("_________");

const fs = require('fs');
const cache = {};
function consistentReadSync(filename) {
if(cache[filename]) {
return cache[filename];
} else {
cache[filename] = fs.readFileSync(filename, 'utf8');
return cache[filename];
}
}

const fs1 = require('fs');
const cache1 = {};
function consistentReadAsync(filename, callback) {
if(cache1[filename]) {
process.nextTick(() => callback(cache1[filename]));
} else {
//асинхронная функция
fs1.readFile(filename, 'utf8', (err, data) => {
cache1[filename] = data;
callback(data);
});
}
}