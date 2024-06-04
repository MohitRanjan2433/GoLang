package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Pointer")
	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer: ", ptr)
	fmt.Println("Value of actual pointer: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", *ptr)
	fmt.Println("New Value is: ", myNumber)

}
