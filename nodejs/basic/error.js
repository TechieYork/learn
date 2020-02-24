'use strict';

var n = 0;
var s;

//Everything is fine
console.log("====== No error ======");
try {
    console.log(n);
} catch (e) {
    if (e instanceof Error) {
        console.log("Error:" + e);
    }
} finally {
    console.log("Finally");
}

//Catched one error
console.log("====== Catched one error ======");
try {
    n = s.length;
} catch (e) {
    if (e instanceof Error) {
        console.log("Error:" + e);
    }
} finally {
    console.log("Finally");
}

//Throw one error
console.log("====== Throw one error ======");
try {
    throw new Error("test throw error");
} catch (e) {
    if (e instanceof Error) {
        console.log("Error:" + e);
    }
} finally {
    console.log("Finally");
}
