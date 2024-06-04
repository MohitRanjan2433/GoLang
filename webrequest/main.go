package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the Web Request...")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("Types of res: %T\n", res)

	data, err := io.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}

	fmt.Println("response: ", string(data));
}