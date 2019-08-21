package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	a = ArrIncr(a)
	fmt.Println(a)

	ArrIncrV2(&a)
	fmt.Println(a)
}

func ArrIncr(arr [3]int) [3]int {
	res := [3]int{}
	for k, v := range arr {
		res[k] = v + 1
	}

	return res
}

func ArrIncrV2(ptr *[3]int) {
	for v := range ptr {
		ptr[v] += 1
	}
}