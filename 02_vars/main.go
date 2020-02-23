package main

import "fmt"

// MAIN TYPES
// string
// bool
// int

func main() {
	// shorthand only inside function
	// name := "Glivol"
	// email:= "glisteningvolume@gmail.com"
	//
	name, email := "Glivol", "glisteningvolume@gmail.com"
	var age int = 29
	const gender = "male"
	var isCool = true
	var float float32 = 2.3

	fmt.Println(name, age, gender, float, email, isCool)
	fmt.Printf("%T\n", email)
}
