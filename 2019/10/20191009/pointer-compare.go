package main

import "fmt"

func main() {
	a := 1

	pointA := &a
	pointB := &a
	fmt.Println(pointA == pointB)

	var pointC *int
	var pointD *int
	fmt.Println(pointC == pointD)
}
