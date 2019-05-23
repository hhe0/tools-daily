package main

import (
	"fmt"
	"math/rand"
)

//与下面的形式等价
//import "fmt"
//import "math"

func main() {
	// 总会返回相同的数字
	fmt.Println("My favorite number is", rand.Intn(10));
}