package main

import "fmt"

func main() {
	x := 10
	changeValueWithoutPointer(x)
	fmt.Println(x)
	changeValueWithPointer(&x)
	fmt.Println(x)
}

func changeValueWithoutPointer(x int) {
	x = 7
}

func changeValueWithPointer(x *int) {
	*x = 7
}
