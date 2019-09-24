package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	s1 = append(s1, 0, 1, 2, 3)
	fmt.Println(s1)

	s2 := make([]int, 0)
	s2 = append(s2, 0, 1, 2, 3)
	fmt.Println(s2)
}
