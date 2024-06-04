package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello Mr.", name)

	//comma - ok || err-ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}
