package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)


	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 934
	// highScores[2] = 734
	// highScores[3] = 534

	// highScores = append(highScores, 955, 766, 321)

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(len(highScores))

	// var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	// fmt.Println(courses)
	// var index int  = 2
	// courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)


	// numbers := []int{1,2,3,4,5}
	// numbers = append(numbers, 8, 10, 12, 14, 16, 20, 25)
	// fmt.Println("Number: ", numbers)
	// fmt.Printf("Number has data type: %T\n", numbers)
	// fmt.Println("Length: ", len(numbers))

	numbers:=make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 55)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length: ", len(numbers))
	fmt.Println("Capacity: ", cap(numbers))

}

