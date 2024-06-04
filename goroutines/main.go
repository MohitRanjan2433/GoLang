package main

import (
	"fmt"
	"time"
)

func sayhello(){
	fmt.Println("Hello world")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("The function is over...")
}

func sayhi(){
	fmt.Println("hiii!!")
}


func main(){
	fmt.Println("learning goroutines")
	go sayhello()
	sayhi()

	time.Sleep(3000 * time.Millisecond)
}