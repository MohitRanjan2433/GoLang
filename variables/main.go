package main

import "fmt"

func main() {
	var username string = "Mohit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	if isLoggedIn == false {
		fmt.Println("Please log in")
	}

	var smallval uint8 = 255
	fmt.Println(smallval);

	var smallfloat float32 = 25.445464551
	fmt.Println(smallfloat)

	var website = "codebox.in"
	fmt.Println(website)
	fmt.Printf("Type of website %T \n", website)

	numberofusers := 56
	fmt.Println(numberofusers)
	fmt.Printf("type of user %T \n", numberofusers)
}
