package main

import "fmt"

func main() {
	//first()
	defer first()
	second()
}

func second() {
	fmt.Println("Second is executed")
}

func first() {
	fmt.Println("First is executed")
}
