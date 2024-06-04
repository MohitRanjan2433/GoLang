package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Litsening on Port 4000...")
}