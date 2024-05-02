//Map test, map started its support on ES6, because {} only support string as key
console.log("====== Map init ======");
var test_map: Map<number, string>[] = new Map([[1, "test1"], [2, "test2"], [3, "test3"]]);
console.log(test_map);

//Set
console.log("====== Set ======");
test_map.set(4, "test4");
console.log(test_map);

//Get
console.log("====== Get ======");
console.log(test_map.get(1));
console.log(test_map.get(5));

//Has
console.log("====== Has ======");
console.log(test_map.has(1));
console.log(test_map.has(5));

//Delete
console.log("====== Delete ======");
test_map.delete(4);
console.log(test_map);

//Iteration for ES5.1 and above
console.log("====== Iteration for ES5.1 and above ======");
test_map.forEach(function (value, key, map) {
    console.log("key:", key, " value:", value, " map:", map);
});

//Iteration for ES6 and above
console.log("====== Iteration for ES6 and above ======");
for (var item of test_map) {
    console.log(item[0], "=>", item[1]);
}
