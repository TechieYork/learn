//Array
console.log("====== Array init ======");
var array: number[] = [1, 2, 3];
var arrayx: Array<number> = [1, 2, 3];

console.log(array);
console.log(arrayx);

//Find value
console.log("====== Find ======");
console.log(array.indexOf(1, 0));

//Some
console.log("====== Some ======");
console.log(array.some(number => number === 1));

//Push or pop back
console.log("====== Push and pop back ======");
array.push(4);
console.log(array);
array.pop();
console.log(array);

//Push or pop front
console.log("====== Push and pop front ======");
array.unshift(0);
console.log(array);
array.shift();
console.log(array);

//Sort
console.log("====== Sort ======");
array.sort(
    function (a, b) {return b - a;}
);
console.log(array);

//Reverse
console.log("====== Reverse ======");
console.log(array.reverse());

//Add up two array
console.log("====== Merge two array ======");
var array1 = [4, 5, 6];

var new_array = array.concat(array1);
console.log(new_array);

//Join to string
console.log("====== Join ======");
console.log(array.join('-'));

//Iteration
console.log("====== Iteration ======");
for (var index = 0; index < array.length; index++)
{
    console.log(array[index]);
}

//Iteration for ES5.1 and above
console.log("====== Iteration for ES5.1 and above ======");
array.forEach(function(item, index, array){
   console.log("index", index, " item:", item, " array:", array);
});

//Iteration for ES6 and above
console.log("====== Iteration for ES6 and above ======");
for (var item of array)
{
    console.log(item);
}

//Every
console.log("====== Find ======");
console.log(array.every(function (value, index, array) {
    if (typeof value === "number")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

console.log(array.every(function (value, index, array) {
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
console.log(array.find(function (value, index, array) {
    if (typeof value === "number")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

console.log(array.find(function (value, index, array) {
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
console.log(array.findIndex(function (value, index, array) {
    if (typeof value === "number")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

console.log(array.findIndex(function (value, index, array) {
    if (typeof value === "string")
    {
        return true;
    }
    else
    {
        return false;
    }
}));

