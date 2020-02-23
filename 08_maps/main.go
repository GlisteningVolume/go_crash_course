package main

import "fmt"

func main() {

	// Define map
	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@email.com"
	emails["Sharon"] = "sharon@email.com"
	emails["Mike"] = "Mike@email.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"]) // real world usage, use unique ID number in case there is more than one bob

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add key values

	eemails := map[string]string{"Bob": "bob@email.com", "Sharon": "ssharon@email.com"}
	eemails["Mike"] = "Mike@email.com"
	fmt.Println(eemails)
}
