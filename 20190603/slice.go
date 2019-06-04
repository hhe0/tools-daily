package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11, 13}

	b := a[:0]
	printSlice(b)

	b = a[:4]
	printSlice(b)

	b = a[2:]
	printSlice(b)

	c := make([]int, 5)
	printSlice(c)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
}