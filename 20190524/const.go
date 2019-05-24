package main

import (
	"fmt"
	"unsafe"
)

const PI float64 = 3.14

const PI2 = 3.14

// 常量表达式中，函数必须是内置函数，否则编译不过
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

// 没有明确赋值的常量与上一个常量相同（不一定指的是值），iota随着常量的定义递增，从0开始
const (
	d = iota
	e
	f = "hello"
	g
	h = 100
	i
	j = iota
)

// 继承了上一个表达式
const (
	l = 1 << iota
	m = 3 << iota
	n
	o
)

func main() {
	fmt.Println(PI, PI2, a, b, c, d, e, f, g, h, i, j, l, m, n, o)
}