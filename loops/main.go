package main

import "fmt"


func main(){
	fmt.Println("Welcom to Loops in go")

	days := []string{"Sunday", "Tuesday", "Fridy", "Sunday"}

	fmt.Println(days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Printf(days[d] + " ")
	// }

	// for i:= range days{
	// 	fmt.Print(days[i] + " ")
	// }

	for index, day := range days{
		fmt.Printf("index is %v and value is %v\n", index, day)
	}
}