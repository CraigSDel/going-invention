package main

import "fmt"

func main() {
	// Immutable object which means that once the objects value is defined it can not be changed
	const pi float64 = 3.141592653589793238

	//multiple variable deceleration
	var (
		varA = 5
		varB = 8
	)

	fmt.Println(varA, varB)

	var name string = "John Doe"
	fmt.Println("The length of the name ", len(name))
}
