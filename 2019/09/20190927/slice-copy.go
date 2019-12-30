package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{1, 2, 3, 4}
	s3 := []int{4, 5, 6}

	n1 := copy(s1, s2)
	fmt.Println(n1, s1, s2)

	n2 := copy(s2, s3)
	fmt.Println(n2, s2, s3)
}
