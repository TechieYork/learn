'use strict';

//Set test, set started its support on ES6
console.log("====== Set init ======");
var test_set: Set<number> = new Set([1, 2, 3]);
console.log(test_set);

//Add
console.log("====== Add ======");
test_set.add(4);
console.log(test_set);

//Has
console.log("====== Has ======");
console.log(test_set.has(1));
console.log(test_set.has(5));

//Delete
console.log("====== Delete ======");
test_set.delete(4);
console.log(test_set);

//Iteration for ES5.1 and above
console.log("====== Iteration for ES5.1 and above ======");
test_set.forEach(function (value, value2, set) {
    console.log("value:", value, " value2:", value2, "set:", set);
});

//Iteration for ES6 and above
console.log("====== Iteration for ES6 and above ======");
for (var item of test_set)
{
    console.log(item);
}
