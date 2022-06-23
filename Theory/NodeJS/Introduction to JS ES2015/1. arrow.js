"use strict";

const numbers = [2, 5, 9, 11, 1];

//function
const even1 = numbers.filter(function(x) {
  return x % 2 === 0;
});
console.log(even1);

//arrow function
const even2 = numbers.filter(x => x % 2 === 0);
console.log(even2);
