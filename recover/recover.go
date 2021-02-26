package main

import "fmt"

func main() {
	fmt.Println(div(3, 0))
	fmt.Println(div(3, 5))
	demoPanic()
}

func div(i int, i2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	return i / i2
}

func demoPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("Panic")
}
