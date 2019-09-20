package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("打印")
	}()

	fmt.Println("Hello World")
}
