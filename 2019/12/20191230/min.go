package main

import "fmt"

func min(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func main() {
	min(1225, 256)

	var s1 []int
	s2 := []int{1, 2, 3, 4}
	s3 := []int{4, 5, 6}

	n1 := copy(s1, s2)
	// 0 [] [1 2 3 4]
	fmt.Println(n1, s1, s2)

	n2 := copy(s2, s3)
	// 3 [4 5 6 4] [4 5 6]
	fmt.Println(n2, s2, s3)
}
