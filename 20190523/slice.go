package main

import (
	"fmt"
)

var (
	slice = []int {1, 2, 4, 8, 16, 32, 64, 128}
)

func main() {
	// 切片的初始化
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	// s[lo:hi] 表示从 lo 到 hi-1 的 slice 元素，含两端
	fmt.Printf("%T\n", p)
	fmt.Println("p[0:1] == ", p[0:2])

	// 省略下标代表从0开始
	fmt.Println("p[:3] == ", p[:3])
	// 省略上标代表到len(s)结束
	fmt.Println("p[3:] == ", p[3:])

	// slice 的零值是nil
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	testNil()

	testAppend()

	testRange()
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func testNil() {
	var z[] int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func testAppend() {
	// TODO: 当新增元素时，其容量的增加具有一定的规律，需要发掘一下~
	var a[] int
	printSlice("a", a)

	a = append(a, 0, 1, 2)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func testRange() {
	// 类似于PHP中的foreach
	for i,v := range slice {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 忽略value
	for i := range slice {
		slice[i] = 1 << uint(i)
	}

	// 忽略key
	for _, value := range slice {
		fmt.Printf("%d\n", value)
	}
}