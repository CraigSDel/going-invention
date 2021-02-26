package main

import "fmt"

func main() {
	// There is only one type of loop in go and that is a for loop
	// (for , while, do while)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// creating a "while" loop
	fmt.Println("Creating a \"while\" loop")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
