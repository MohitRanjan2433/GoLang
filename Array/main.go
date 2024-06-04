package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	// var fruitList [4]string
	// fruitList[0] = "Apple"
	// fruitList[1] = "Mango"
	// fruitList[3] = "Peach"

	// fmt.Println("Fruit list is: ", fruitList)
	// fmt.Println("Fruit list is: ", len(fruitList))

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[0:3])
	fmt.Println(fruitList)

}
