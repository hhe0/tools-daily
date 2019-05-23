package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
}

func add(x int, y int) int {
	return x + y
}