package main

import "fmt"

func main() {
	ids := []int{33, 23, 14, 64, 1, 2} // Slice

	// Loop through ids

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids { // underscore used to replace variable that will not be called to prevent error
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map

	eemails := map[string]string{"Bob": "bob@email.com", "Sharon": "ssharon@email.com"}

	for k, v := range eemails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range eemails {
		fmt.Println("Name: " + k)
	}
}
