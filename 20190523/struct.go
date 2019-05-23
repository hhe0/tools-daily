package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	// 忽略其中的一个字段，会取默认值
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	v4 = &Vertex{1, 2}
)

func main() {
	// 结构体变量的初始化
	v := Vertex{1, 2}
	// 结构体赋值
	v.X = 4
	fmt.Println(v.X)

	// 结构体指针
	p := &v
	p.X = 111
	fmt.Println(v.X)

	fmt.Println(v1, v2, v3, *v4)
}