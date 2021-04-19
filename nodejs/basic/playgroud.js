console.log("====== playground ======");

var fn = function foo() {};
fn.name = "asdf";

var obj = {};

console.log(obj.prototype);
console.log(fn.prototype.name);
