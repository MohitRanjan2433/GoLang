package main

import (
	"fmt"
	"net/url"
)

// const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456hb"
// const myurl string = "https://www.javatpoint.com/reactjs-tutorial"
const myurl string = "https://www.reddit.com/r/cpp_questions/comments/updx25/need_help_installing_latest_g_on_windows/?rdt=52714"

func main() {

	fmt.Println("Welcome to heading urls in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The type of query params are %T\n", qparams)

	// fmt.Println(qparams["coursename"])

	// for i, val := range qparams{
	// 	fmt.Println("Params of: ", i, val)
	// }

	partsofUrl := &url.URL{
		Scheme:"https",
		Host: "www.javapoint.com",
		Path: "/reinforcement-learning",
	}

	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)
}