package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in golang")

	Mohit := User{"Mohit", "mohit.dev@gmail.com", true, 20}
	fmt.Println(Mohit)

	fmt.Printf("Mohit Details are %+v\n", Mohit)
	fmt.Println("Name is %v and email is %v and age is %v", Mohit.Name, Mohit.Email, Mohit.Age)
	Mohit.GetStatus()
	Mohit.NewMail()


}

type User struct{
	Name string
	Email string
	Status bool
	Age int

}


func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is: ", u.Email)
}