'use strict';

function abs(number)
{
    if (typeof number !== 'number')
    {
        throw 'Not a number';
    }

    if (number >= 0)
    {
        return number;
    }
    else
    {
        return -number;
    }
}

//Test abs
console.log("====== abs() ======");
console.log(abs(-1));
console.log(abs(2));
console.log(abs(-3, -4, 5));

try
{
    console.log(abs());
}
catch (e)
{
    console.log(e);
}

try
{
    console.log(abs("asdf"));
}
catch (e)
{
    console.log(e);
}

//Test argument
console.log("====== arguments ======");

function testArgument(a, b)
{
    for (let index in arguments) {
        console.log(index);
    }

    for (let argu of arguments) {
        console.log(argu);
    }
}

testArgument("argu1", "argu2", "argu3", "argu4");

//Test rest
console.log("====== rest ======");

function testRest(a, b, ...rest) {
    for (let index in arguments) {
        console.log(index);
    }

    for (let argu of arguments) {
        console.log(argu);
    }

    for (let index in rest) {
        console.log(index);
    }

    for (let r of rest) {
        console.log(r);
    }
}

testRest("argu1", "argu2", "rest1", "rest2");

