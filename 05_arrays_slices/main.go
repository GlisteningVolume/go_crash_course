package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and Assign at the same time
	fruitArr2 := [2]string{"Papaya", "Banana"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(fruitArr2)

	fruitSlice := []string{"Apple", "Orange", "Lemon", "Lime"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice)) // find count of array
	fmt.Println(fruitSlice[1:4]) // find array elements in range
}
