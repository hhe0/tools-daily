package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := s[2:4]
	fmt.Println(s1, len(s1), cap(s1))

	s2 := s[2:4:8]
	fmt.Println(s2, len(s2), cap(s2))
}
