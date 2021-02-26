package main

import "fmt"

func main() {
	const pi float64 = 3.141592653589793238
	var a int = 5
	var b float32 = 3.14

	// multiple variables
	var (
		c = 8
		d = 7
	)

	x, y := 15, 8

	fmt.Println("PI : ", pi)
	fmt.Println("A : ", a)
	fmt.Println("B : ", b)
	fmt.Println("C , D : ", c, d)
	fmt.Println("X , Y : ", x, y)
	fmt.Println("-------------------------")
	fmt.Println("x + y : ", x+y)
	fmt.Println("x - y : ", x-y)
	fmt.Println("x * y : ", x*y)
	fmt.Println("x / y : ", x/y)
	fmt.Println("x % y : ", x%y)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
