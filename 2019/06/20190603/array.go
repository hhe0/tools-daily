package main

import "fmt"

func main() {
	var a [4]int
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	b := [4]int{0, 1, 2, 3}
	for j := 0; j < len(b); j++ {
		fmt.Println(b[j])
	}

	var s = b[1:3]
	for k := 0; k < len(s); k++ {
		fmt.Println(s[k])
	}

	t := []bool{true, true, false, false}
	for l := 0; l < len(t); l++ {
		fmt.Println(t[l])
	}

	var z []int
	fmt.Println(z, len(z), cap(z))
}
