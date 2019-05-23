package main

import (
	"fmt"
)

func main() {
	// 数组的长度是其类型的一部分，因此数组不能改变大小
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Printf("%T\n", a)
}
