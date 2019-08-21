package main

import "fmt"

func main() {
	a := 0
	if true {
		a, b := test(a)
		fmt.Println(a, b)
	}
	fmt.Println(a)
}

func test(a int) (int, int) {
	return a + 1, 0
}
