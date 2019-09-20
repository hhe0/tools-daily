package main

import "fmt"

func main() {
	if err := recover(); err != nil {
		fmt.Println("处理异常")
	}

	panic("触发异常")

	fmt.Println("打印")
}
