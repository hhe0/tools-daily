package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}

	for i, ele := range arr1 {
		fmt.Printf("Array at index %d is %d\n", i, ele)
	}

	// 切片
	a := [...]string{"a", "b", "c", "d"}
	fmt.Printf("%T\n", a)
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	b := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", b)
	for i := range b {
		fmt.Println("Array item", i, "is", b[i])
	}

	arr2 := new([5]int)
	fmt.Printf("%T\n", arr2)
}
