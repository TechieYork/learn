'use strict';

class Person {
    constructor(name, age) {
        this.name = name;
        this.age = age;
    }

    sayHello() {
        console.log(`Hello I am ${this.name} and I am ${this.age} years old!`);
    }
}

module.exports.Person = Person;