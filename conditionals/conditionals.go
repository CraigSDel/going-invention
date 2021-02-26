package main

import "fmt"

func main() {
	var age int = 18
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}

	switch age {
	case 16:
		fmt.Println("You can get a motorbike licenses")
	case 18:
		fmt.Println("You can get a car licenses")
	default:
		fmt.Println("Happy walking")
	}

}
