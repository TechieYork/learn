'use strict';

//Class
console.log("====== Class ======");
var Person =
    {
        name: "",
        age: 0,
        sayHello: function () {
            console.log("Hello " + this.name);
        }
    };

//Create a class
console.log("====== Create class with a custom function======");
function createPerson(name) {
    var p = Object.create(Person);

    p.name = name;
    return p;
}

var test_person = createPerson("Foo");
test_person.sayHello();

//Create a class with constructor
console.log("====== Create class with a constructor ======");
function Student(name) {
    this.name = name;
    this.age = 0;
}

//This function will make all instances share one function, otherwise each instance will have its own sayHello function
//And that will cause memory waste
Student.prototype.sayHello = function () {
    console.log("Hello " + this.name);
};

var test_student = new Student("Bar");  //new must to be added(this will returned), otherwise it will return the function not a object
test_student.sayHello();

//Create a class for ES6 and above
console.log("====== Create class for ES6 and above ======");
class Teacher {
    constructor (name) {
        this.name = name;
    }

    sayHello() {
        console.log("Hello " + this.name);
    }
}

var test_teacher = new Teacher("Wola");
test_teacher.sayHello();

//Create a subclass for ES6 and above
console.log("====== Create a subclass for ES6 and above");
class SportTearcher extends Teacher {
    constructor (name, sport) {
        super(name);
        this.sport = sport;
    }

    whatSportToTeach() {
        console.log("I teach " + this.sport);
    }
}

var test_sport_teacher = new SportTearcher("William", "Football");
test_sport_teacher.sayHello();
test_sport_teacher.whatSportToTeach();

