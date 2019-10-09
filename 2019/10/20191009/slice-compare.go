package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 0}
	sliceB := []int{1, 2}
	fmt.Print(sliceA == sliceB)
}
