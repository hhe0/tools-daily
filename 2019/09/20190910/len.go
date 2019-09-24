package main

import "fmt"

func main() {
	str1 := "hello world"
	fmt.Println(len(str1))
	fmt.Println(len([]rune(str1)))

	str2 := "你好 世界"
	fmt.Println(len(str2))
	fmt.Println(len([]rune(str2)))
}
