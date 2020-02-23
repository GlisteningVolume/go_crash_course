package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

//Greeting Method

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

// hasBirthday method to (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getsMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Init person using struct

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// Alternative
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName, person1.age) // 25
	person1.age++
	fmt.Println(person1.firstName, person1.age) // 26
	person1.age = person1.age + 2
	fmt.Println(person1.firstName, person1.age) // 28

	fmt.Println(person1.greet()) // Hello, my name is Samantha Smith and i am 28

	person1.hasBirthday()

	fmt.Println(person1.greet()) // Hello, my name is Samantha Smith and i am 29

	person1.getsMarried("Williams")

	fmt.Println(person1.greet())

	person2.getsMarried("Thompson")

	fmt.Println(person2.greet())

}
