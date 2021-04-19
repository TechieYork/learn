'use strict';

//Variable
console.log("====== Variable ======");
var test_number = 1;
var test_float = 1.11;
var test_string = 'hello';

console.log(test_number, test_float, test_string);

const test_const = 1; //For ES6 and above, has block level scope
console.log(test_const);

//Type
console.log("====== Type ======");
console.log(typeof test_number);
console.log(typeof test_float);
console.log(typeof test_string);

//Comparison
console.log("====== Comparison ======");
// == and ===, === compares both the value and type
var a = 1;
var b = '1';
var c = 1;

//null, NaN, undefined, 0 and '' will be take as false in comparison
if (!null && !NaN && !undefined && !0 && !'')
{
    console.log("null!");
}
if (!NaN)
{
    console.log("NaN!");
}
if(!undefined)
{
    console.log("undefined!");
}
if (!0)
{
    console.log("0!");
}
if (!'')
{
    console.log("''!");
}


console.log(a == b);
console.log(a === b);
console.log(a == c);
console.log(a === c);

//Object
console.log("====== Object ======");
var test_person =
    {
        name: "bob",
        age: 20,
        zip: null,
        tel_num: 987654321,
        'middle-school': "river high",    //Special key name, use "" to wrap it, and could only use ['middle-school'] to use it rather than .
        getAge: function () {
            return this.age;
        }
    };

console.log(test_person);
console.log(test_person["middle-school"]);
console.log(test_person.tel_num);
console.log(test_person.not_exist); //If the property not exist, undefined will be returned

//Test function
console.log(test_person.getAge());

delete test_person.age;  //Delete a variable
console.log(test_person);

//Test if a variable is in the class
console.log('name' in test_person);
console.log('age' in test_person);

//Because all objects inherited from object so the property will inherited such as 'toString' in test_person
//In this way, if the only thing we wanna do is check whether test_person has this property then use hasOwnProperty
console.log(test_person.hasOwnProperty('name'));
console.log(test_person.hasOwnProperty('age'));

//var and let(for ES6 and above)
//They have different variable scope, var for function and let for block
console.log("====== var and let ======");
function testVarLet()
{
    //Both v and i has functional level scope
    var v = 1;
    for (var i = 1; i < 3; i++)
    {
        console.log(i);
    }

    console.log(i);

    //J has block level scope
    for (let j = 1; j < 3; j++)
    {
        console.log(j);
    }

    //console.log(j); //Error not defined
}

testVarLet();

//map, filter and reduce
console.log("====== map, filter and reduce ======");

var filtered = ["1", "2", "3"].map((value, index, array) => {
    // map, deal with each item in the array and return new array
    // if return nothing or without return statement, new array will be filled with undefined
    console.log(value);
    return value;
}).filter((value, i, self) => {
    // filter, deal with each item in the array and return new array
    // if return value is true, then the value will be added to the new array, otherwise, it will be discarded
    console.log([i], value, self);
    return true;
}).reduce((previousValue, currentValue, currentIndex) => {
    // reduce, deal with each item in the array and return a new item
    // such as, calculate the number, connect the strings
    return previousValue + currentValue;
}, "");

console.log(filtered);

