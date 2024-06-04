package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main() {
	fmt.Println("Welcome to the Get request in Golang")
	performGetRequest()
	
}

func performGetRequest(){
	const myurl = "http://localhost:4000/candidate/vote/count/all"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content,_ := ioutil.ReadAll(response.Body)
	byteCount,_ := responseString.Write(content)

	fmt.Println("ByteCount is: ",byteCount)
	fmt.Println(responseString.String())
	

	// fmt.Println(string(content))
}