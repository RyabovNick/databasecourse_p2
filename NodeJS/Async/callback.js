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
