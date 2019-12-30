package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Println(s)

	m := make(map[string]interface{}, 0)
	fmt.Println(m)

	c := make(chan bool, 1)
	fmt.Println(c)
}
