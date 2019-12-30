package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
}

func (c Circle) Area(radius float32) float32 {
	return math.Pi * radius * radius
}

type Shapes struct {
	circle Circle
	//rectangle Rectangle
}

func main() {
	var shapes Shapes
	shapes.circle = Circle{2}
	//shapes.Rectangle = Rectangle{1, 2}

	fmt.Println(shapes.circle)
	//fmt.Println(shapes.Rectangle)
}