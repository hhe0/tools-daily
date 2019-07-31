package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Minc"
	fmt.Println("============")
	for i := 0; i < len(name); i++ {
		fmt.Println(string(name[i]), reflect.TypeOf(name[i]))
	}
	fmt.Println("============")
	for _, n := range name {
		fmt.Println(string(n), reflect.TypeOf(n))
	}

	chineseName := "何宏"
	fmt.Println("============")
	for i := 0; i < len(chineseName); i++ {
		fmt.Println(string(chineseName[i]), reflect.TypeOf(chineseName[i]))
	}
	fmt.Println("============")
	for _, n := range chineseName {
		fmt.Println(string(n), reflect.TypeOf(n))
	}
	fmt.Println("============")
}
