package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 总会返回相同的数字
	fmt.Println("My favorite number is", rand.Intn(10));
}