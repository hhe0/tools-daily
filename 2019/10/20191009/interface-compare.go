package main

import "fmt"

func main() {
	var interfaceA interface{}
	var interfaceB interface{}
	interfaceA = 1
	interfaceB = 1
	fmt.Println(interfaceA == interfaceB)

	var interfaceC interface{}
	var interfaceD interface{}
	interfaceC = 2
	interfaceD = "2"
	fmt.Println(interfaceC == interfaceD)

	var interfaceE interface{}
	var interfaceF interface{}
	fmt.Println(interfaceE == interfaceF)
}
