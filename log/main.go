// greetings.go

package main

import (
	"errors"
	"fmt"
)

// Hello returns a greeting message.
func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty name")
	}
	message := fmt.Sprintf("Hello, %s!", name)
	return message, nil
}


func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := Hello("John")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
