package main

import "fmt"

var c, python, java int = 1, 2, 3

func main() {
	var i int
	// 短声明变量只能用于函数内部
	k := 4
	// 0 false false false
	fmt.Println(i, c, python ,java, k)
}