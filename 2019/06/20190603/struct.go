package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	fmt.Println(a.X, a.Y)

	p := &a
	fmt.Println(a.X, p.Y)

	b := Vertex{X:1}
	fmt.Println(b.X, b.Y)

	c := Vertex{}
	fmt.Println(c.X, c.Y)

	d := &Vertex{Y:1}
	fmt.Println(d.X, d.Y)

	var e Vertex
	fmt.Println(e.X, e.Y)
}
