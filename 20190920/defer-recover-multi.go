package main

import "fmt"

func main() {
	deferCallRecover()
}

func deferCallRecover() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
		if err := recover(); err != nil {
			fmt.Println("打印中捕捉异常")
		}
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")

	fmt.Println("Hello World")
}
