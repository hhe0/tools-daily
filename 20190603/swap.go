package main

import "fmt"

func main() {
	i := 1
	j := 2
	i, j = swap(i, j)
	fmt.Println(i, j)
}

func swap(x, y int) (int, int) {
	return y, x
}

