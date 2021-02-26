package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle := Rectangle{30, 40}
	circle := Circle{2}
	fmt.Println("Height ", rectangle.height, " Width ", rectangle.width)
	fmt.Println("Rectangle area ", rectangle.area())
	fmt.Println("Rectangle area ", getArea(rectangle))
	fmt.Println("Circle area ", circle.area())
	fmt.Println("Circle area ", getArea(circle))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}
