package main

import "fmt"

func main() {

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Wtch Out"
	} else {
		result = "Exactly 10 login Count"
	}

	fmt.Println(result)

	if num:=3; num < 10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is greater than 10")
	}

}