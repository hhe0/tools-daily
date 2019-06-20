package main

import "fmt"

type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s Square)Area() float32 {
	return s.side * s.side
}

type Rectangle struct {
	width float32
	height float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func main() {
	s := Square{2}
	r := Rectangle{1, 3}

	shapes := []Shape{s, r}
	for n := range shapes {
		// 检测类型
		if _, ok := shapes[n].(*Rectangle); ok {
			fmt.Print(shapes[n])
		}
		fmt.Println("shape detail: ", shapes[n])
		fmt.Println("shape area: ", shapes[n].Area())
	}
}