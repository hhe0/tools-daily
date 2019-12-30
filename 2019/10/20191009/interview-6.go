package main

import "fmt"

func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)

	b := make([]int, 1, 4)
	fmt.Printf("%+v\n", b)
	ap(b)
	fmt.Printf("%+v\n", b)
	app(b)
	fmt.Printf("%+v\n", b)
}

func ap(a []int) {
	a = append(a, 10)
	fmt.Println(len(a), cap(a))
}

func app(a []int) {
	a[0] = 1
}
