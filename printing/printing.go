package main

import "fmt"

func main() {
	const pi float64 = 3.141592653589793238
	x := 5
	isbool := true

	// printing a float number
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)
	// Shows the type
	fmt.Printf("%T \n", isbool)
	fmt.Printf("%t \n", pi)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 25)
	fmt.Printf("%c \n", 34)
	fmt.Printf("%x \n", 15)
	//scientific notation
	fmt.Printf("%e \n", pi)

	fmt.Println("")
	var name string = "John Doe"
	fmt.Println(name)
	fmt.Println(len(name))
}
