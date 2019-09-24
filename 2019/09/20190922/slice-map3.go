package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key := range slice {
		m[key] = &slice[key]
	}

	for k, v := range m {
		fmt.Println(k, ">", *v)
	}
}
