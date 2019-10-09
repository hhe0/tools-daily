package main

import "fmt"

func main() {
	var chanC chan int
	var chanD chan int
	fmt.Println(chanC == chanD)

	chanA := make(chan int)
	chanB := make(chan int)
	fmt.Println(chanA == chanB)

	chanC = chanA
	chanD = chanA
	fmt.Println(chanC == chanD)
}
