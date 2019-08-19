'use strict';

const Person = require('./person/person');
const Teacher = require('./person/teacher');

var p = new Person.Person("foo", 20);
p.sayHello();

var t = new Teacher.Teacher("bar", 30, "English");
t.sayHello();

var s = new Teacher.SportTeacher("larry", 40, "Basketball");
s.sayHello();