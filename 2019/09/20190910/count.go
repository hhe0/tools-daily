package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"

	fmt.Println(strings.Count(str, "h"))
	fmt.Println(strings.Count(str, "l"))
	fmt.Println(strings.Count(str, "0"))
}
