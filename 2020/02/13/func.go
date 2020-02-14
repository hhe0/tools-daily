package main

import "fmt"

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

	nums1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(nums1...))

	//nums2 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	//fmt.Println(sum())
}

func sum(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}
