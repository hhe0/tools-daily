package main

import "fmt"

type MyIntA int
type MyIntB = int

func main() {
	intA := 1

	var myIntA MyIntA = intA
	//var myIntA MyIntA = MyIntA(intA)
	fmt.Println(myIntA)

	var myIntB MyIntB = intA
	fmt.Println(myIntB)
}
