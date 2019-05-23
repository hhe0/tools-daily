package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(7))
}

func add(x int, y int) int {
	return x + y
}

// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
func sub(x, y int) int {
	return x - y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 没有参数的 return 语句返回结果的当前值。也就是`直接`返回
// 直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

