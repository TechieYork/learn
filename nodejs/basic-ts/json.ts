//Json
console.log("====== JSON ======");
var test_person =
{
    name: "bob",
    age: 20,
    zip: null,
    tel_num: 987654321,
    'middle-school': "river high",    //Special key name, use "" to wrap it, and could only use ['middle-school'] to use it rather than .
};

//Marshal
console.log("====== Marshal ======");
var jsonStr = JSON.stringify(test_person);
console.log(jsonStr);

//Marshal with specific keys and format
console.log("====== Marshal with specific keys and format ======");
console.log(JSON.stringify(test_person, ["name", "age"], "    "));

//Marshal with functions
console.log("====== Marshal with functions ======");
console.log(JSON.stringify(test_person, function (key, value) {
    console.log("key:" + key + ", value:" + value);
    if (typeof value === "string") {
        return value.toUpperCase();
    }
    return value;
}, "    "));

//Unmarshal
console.log("====== Unmarshal ======");
var test_person2 = JSON.parse(jsonStr);
console.log(test_person2);


