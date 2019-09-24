package main

import "fmt"

func main() {
	s := make([]int, 2, 5)
	s = append(s, 0, 1, 2, 3)
	fmt.Println(s)
}
