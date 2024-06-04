package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("We are learning JSON")

	person := Person{Name: "John", Age: 34, IsAdult: true}

	//Convert person into JSON Encoding(Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	fmt.Println("Person Data is: ", string(jsonData))

	//Decoding (Unmarshalling)
	var personData Person

	err = json.Unmarshal(jsonData, &personData)
	if err != nil{
		panic(err)
	}

	fmt.Println("Person Data is:", personData)
}