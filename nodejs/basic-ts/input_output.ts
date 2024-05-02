const readline = require('readline');

//Input
var rl = readline.createInterface({
   input: process.stdin,
   output: process.stdout
});

rl.question("How old are you?\r\n", function (answer) {
    console.log("I am " + answer + " years old!");
});

rl.on("close", function () {
    console.log("Process exit!");
});
