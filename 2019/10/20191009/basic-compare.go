package main

import "fmt"

func main() {
	boolA := true
	boolB := true
	fmt.Println(boolA == boolB)

	intA := 1
	intB := 1
	fmt.Println(intA == intB)

	floatA := 3.14
	floatB := 3.14
	fmt.Println(floatA == floatB)

	complexA := 1 + 1i
	complexB := 1 + 1i
	fmt.Println(complexA == complexB)

	stringA := "string"
	stringB := "string"
	fmt.Println(stringA == stringB)
}
