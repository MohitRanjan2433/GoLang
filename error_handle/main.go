package main

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	message := fmt.Sprintf("Hello, %s!", name)
	return message, nil
}

func main() {
	greeting, err := Hello("John")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(greeting)
	}

	// Example with empty name
	greeting, err = Hello("")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(greeting)
	}
}
