package main

import "fmt"

func main() {
	shortDeceleration()
	longDeceleration()

	numSlice := []int{5, 4, 3, 2, 1}
	sliced := numSlice[3:5]
	fmt.Println(sliced)
	slicedToAPoint := numSlice[:5]
	fmt.Println(slicedToAPoint)
	copy(sliced, slicedToAPoint)
	fmt.Println(sliced)
	fmt.Println(slicedToAPoint)

	sliceAppended := append(numSlice, 9, 8, 7)
	fmt.Println(sliceAppended)
}

func shortDeceleration() {
	EvenNum := [5]int{0, 2, 4, 6, 8}
	fmt.Println(EvenNum[2])

	for _, value := range EvenNum {
		fmt.Println(value)
	}

}

func longDeceleration() {
	var EvenNum [5]int
	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println(EvenNum[2])

	for _, value := range EvenNum {
		fmt.Println(value)
	}
}
