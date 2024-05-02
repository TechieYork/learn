//Create an interface
interface IHello {
    sayHello(): void;
}

sayHello = function(sh: IHello) {
    sh.sayHello();
}


//Create a class
class Person implements IHello {
    constructor (name) {
        this.name = name;
    }

    sayHello() {
        console.log("Hello " + this.name);
    }
}

//Create a subclass
class Student extends Person {
    constructor (name, age) {
        super(name);
        this.age = age;
    }

    sayHello() {
        console.log("Hello " + this.name + ", your age is " + this.age);
    }
}

//Create a subclass
class Teacher extends Person {
    constructor (name, subject) {
        super(name);
        this.subject = subject;
    }

    sayHello() {
        console.log("Hello " + this.name + ", your subject is " + this.subject);
    }
}

//Create an object
var person = new Person("John");
person.sayHello();
var stu = new Student("Tom", 20);
stu.sayHello();
var teacher = new Teacher("Jane", "Math");
teacher.sayHello();

sayHello(person);
sayHello(stu);
sayHello(teacher);
