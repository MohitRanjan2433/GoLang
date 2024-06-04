package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in golang")

	Mohit := User{"Mohit", "mohit.dev@gmail.com", true, 20}
	fmt.Println(Mohit)

	fmt.Printf("Mohit Details are %+v\n", Mohit)
	fmt.Printf("Name is %v and email is %v and age is %v", Mohit.Name, Mohit.Email, Mohit.Age)


}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}