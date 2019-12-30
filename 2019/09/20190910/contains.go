package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	fmt.Println(strings.Contains(str, "hello"))
	fmt.Println(strings.Contains(str, "hell0"))
	fmt.Println(strings.Contains(str, ""))
}
