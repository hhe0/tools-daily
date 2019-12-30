package main

import "fmt"

type People struct {
	Sex   int
	Name  string
	Alive bool
}

func main() {
	p := new(People)

	fmt.Println(p)
	fmt.Println(*p)
}
