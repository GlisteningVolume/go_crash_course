package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // output is *int which stands for pointer

	// Use * to read val from address
	fmt.Println(b)  // address where value is stored
	fmt.Println(*b) // value

	// Change value with pointer
	fmt.Println(a) // 5
	*b = 10
	fmt.Println(a) // 10
}
