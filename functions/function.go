package main

import "fmt"

func main() {
	x, y := 5, 6
	fmt.Println(add(x, y))
	fmt.Println(factorial(5))
}

//demonstrating recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func add(x int, y int) interface{} {
	return x + y
}
