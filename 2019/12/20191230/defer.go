package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("defer 21")
		}
		fmt.Println("defer 22")
	}()

	printDefer()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("defer 3")
		}
	}()

	panic("Error Occured!")

	fmt.Println("这是一个正常的语句")
}

func printDefer() {
	defer func() {
		fmt.Println("defer 4")
	}()
}
