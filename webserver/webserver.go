package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handler2)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to my server")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
}
