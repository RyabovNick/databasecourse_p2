"use strict";

class Person {
  constructor(name, surname, age) {
    this.name = name;
    this.surname = surname;
    this.age = age;
  }

  getFullName() {
    return this.name + " " + this.surname;
  }

  static older(person1, person2) {
    return person1.age >= person2.age ? person1 : person2;
  }
}

class PersonWithMiddleName extends Person {
  constructor(name, middlename, surname, age) {
    super(name, surname, age);
    this.middlename = middlename;
  }

  getFullName() {
    return this.name + " " + this.middlename + " " + this.surname;
  }
}

var per = new Person("Ivan", "Ivanov", 23);

console.log(per.getFullName());

var perwmn = new PersonWithMiddleName("Maria", "MName", "Mar", 22);

console.log(perwmn.getFullName());

console.log(Person.older(per, perwmn));
