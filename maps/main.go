package main

import (
	"fmt"
	// "golang.org/x/text/language"
)

func main() {
	fmt.Println("Welcome to Maps in Golang")

	// Declare a map with string keys and values
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of All Languages: ", languages)
	fmt.Println("JS Short for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of All languages", languages)

	//loops in golang

	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key,value)
	}

	for _, value := range languages{
		fmt.Printf("For key v, value is %v\n",value)
	}

}
