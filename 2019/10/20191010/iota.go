package main

import "fmt"

const (
	a, b = iota * 1, iota * 2
	c, d
	e, f
)

const (
	a1 = iota
	b1 = 233
	c1
	d1 = iota
	e1
)

func main() {
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)

	fmt.Println(a1, b1, c1, d1, e1)
}
