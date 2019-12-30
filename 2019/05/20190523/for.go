package main

import "fmt"

func main() {
	sum := 0
	// 前置、后置语句可以为空
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	res := 1
	// 类似于else的功能
	for res < 1000 {
		res += res
	}
	fmt.Println(res)
}