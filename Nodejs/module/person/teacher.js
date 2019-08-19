'use strict';

const Person = require('./person');

class Teacher extends Person.Person {
    constructor(name, age, subject) {
        super(name, age);
        this.subject = subject;
    }

    sayHello() {
        console.log(`Hello I am ${this.name} and I am ${this.age} years old! I teach ${this.subject}`);
    }
}

class SportTeacher extends Teacher {
    constructor(name, age, subject) {
        super(name, age, subject);
    }
}

module.exports.Teacher = Teacher;
module.exports.SportTeacher = SportTeacher;