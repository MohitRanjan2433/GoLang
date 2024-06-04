package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Post Req in Go")

	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:4000/candidate"

	requestBody := strings.NewReader(`
		{
			"name": "Rajas",
			"party": "INP",
			"age": 32,
			"voteCount": 5
		}
	`)

	// Create a new POST request
	req, err := http.NewRequest("POST", myurl, requestBody)
	if err != nil {
		panic(err)
	}

	// Set content type and authorization headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "abcgbssbshYJUJYJjssns") // Replace with your token

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
