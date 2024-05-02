//Async in order
function testLog(duration, text) {
    setTimeout(function() {
        console.log("Log printed after " + duration + "ms, text is:" + text);
    }, duration);
    //Atomics.wait(new Int32Array(new SharedArrayBuffer(4)), 0, 0, 1000);
    //console.log("Log printed after " + duration + "ms, text is:" + text);
}

var aysncLog = async function () {
    await testLog(2000, "test 1");
    console.log("test 1 done");
    await testLog(1000, "test 2");
    console.log("test 2 done");
};

aysncLog().then(function(result){
    console.log(`test 1 & 2 done, result: ${result}`);
});

//Async simuteniously
var asyncTestLog1 = async function () {
    await testLog(2000, "test async 1");
};

var asyncTestLog2 = async function () {
    await testLog(1000, "test async 2");
};

asyncTestLog1();
asyncTestLog2();
