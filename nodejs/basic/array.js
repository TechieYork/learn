'use strict';

//Array
console.log("====== Array init ======");
var test_array = [1, 2, 3];

console.log(test_array);

//Find value
console.log("====== Find ======");
console.log(test_array.indexOf(1, 0));

//Some
console.log("====== Some ======");
console.log(test_array.some(number => number === 1));

//Push or pop back
console.log("====== Push and pop back ======");
test_array.push(4);
console.log(test_array);
test_array.pop();
console.log(test_array);

//Push or pop front
console.log("====== Push and pop front ======");
test_array.unshift(0);
console.log(test_array);
test_array.shift();
console.log(test_array);

//Sort
console.log("====== Sort ======");
test_array.sort(
    function (a, b) {return b - a;}
);
console.log(test_array);

//Reverse
console.log("====== Reverse ======");
console.log(test_array.reverse());

//Add up two array
console.log("====== Merge two array ======");
var test_array1 = [4, 5, 6];

var new_array = test_array.concat(test_array1);
console.log(new_array);

//Join to string
console.log("====== Join ======");
console.log(test_array.join('-'));

//Iteration
console.log("====== Iteration ======");
for (var index = 0; index < test_array.length; index++)
{
    console.log(test_array[index]);
}

//Iteration for ES5.1 and above
console.log("====== Iteration for ES5.1 and above ======");
test_array.forEach(function(item, index, array){
   console.log("index", index, " item:", item, " array:", array);
});

//Iteration for ES6 and above
console.log("====== Iteration for ES6 and above ======");
for (var item of test_array)
{
    console.log(item);
}

//Every
console.log("====== Find ======");
console.log(test_array.every(function (value, index, array) {
    if (typeof value === "number")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

console.log(test_array.every(function (value, index, array) {
    if (typeof value === "string")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

//Find
console.log("====== Find ======");
console.log(test_array.find(function (value, index, array) {
    if (typeof value === "number")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

console.log(test_array.find(function (value, index, array) {
    if (typeof value === "string")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

//Find index
console.log("====== Find index ======");
console.log(test_array.findIndex(function (value, index, array) {
    if (typeof value === "number")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

console.log(test_array.findIndex(function (value, index, array) {
    if (typeof value === "string")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

