package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct{
	UserID int `json:"userId"`
	Id int `json:"id"`
	Title string `"json:string"`
	Completed bool `json:completed"`
}

func performGetRequest(){
	res, err := http.Get("https://jsonplaceholder.typicode.com/photos/1")
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Response:", string(data))
	// fmt.Println("Length of Data: ", len(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil{
		panic(nil)
	}
	todo.Completed = true
	fmt.Println("Todo: ", todo)
	fmt.Println("Todo Title: ", todo.Title)
	fmt.Println("Completed Response: ", todo.Completed)
}

func performPostRequest(){
	todo := Todo {
		UserID: 23,
		Title: "Mohit Ranjan",
		Completed: true,
	}

	//Conver the Todo struct in JSON
	jsonData, err := json.Marshal(todo)
	if err != nil{
		panic(err)
	}

	//convert json data to string
	jsonString := string(jsonData);

	//convert string dta to reader
	jsonReader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos"

	//send POST request
	res, err := http.Post(myurl, "application/json", jsonReader)

	if err != nil{
		panic(err)
	}

	defer res.Body.Close()
	 
	data, err := ioutil.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}

	fmt.Println("Response: ", string(data))
	fmt.Println("Successfully Data Posted")
}

func performUpdateRequest(){
	todo := Todo {
		UserID: 35,
		Title: "Mohit Raj Golang",
		Completed: true,
	}

	//convert the Todo Struct to JSON
	jsonData, err := json.Marshal(todo)

	if err != nil{
		panic(err)
	}

	//convert json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT request...

	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)

	if err != nil{
		panic(err)
	}

	req.Header.Set("Content-type", "application/json")

	//send the req
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
		fmt.Println("Err", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	fmt.Println("Response: ", string(data))
	fmt.Println("Status: ", res.Status)
}

func performDeleteMethod(){
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create delete request

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil{
		panic(err)
	}

	//send the request
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Response status: ", res.Status)
}

func main(){
	fmt.Println("Welcome to the post request...")

	// performGetRequest()

	// performPostRequest()

	// performUpdateRequest()

	performDeleteMethod()
}