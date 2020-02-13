package main

import "fmt"

func main() {
	ages := map[int]string{1: "1", 2: "2"}
	fmt.Println(ages[1])

	if age, ok := ages[3]; ok {
		fmt.Println(age)
	}
}
