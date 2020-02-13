package samplePackage

import "fmt"

var a [10]int

func init() {
	fmt.Println(1111)

	PrintA()
}

func PrintA() {
	for i := range a {
		fmt.Println(i)
	}
}
