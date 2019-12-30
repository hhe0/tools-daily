package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"

	splitStr := strings.Split(str, "l")
	fmt.Println(splitStr)
}
