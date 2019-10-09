package main

import "fmt"

func main() {
	fmt.Print(Add == Multi)
}

func Add(a int, b int) int {
	return a + b
}

func Multi(a int, b int) int {
	return a + b
}
