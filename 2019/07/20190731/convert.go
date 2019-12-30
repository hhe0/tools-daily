package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := "24"
	// cannot convert expression of type string to type int
	//fmt.Println(int(age))

	res, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	name := 17
	fmt.Println(string(name))

	// 没有 err
	r := strconv.Itoa(name)
	fmt.Println(r)
}
