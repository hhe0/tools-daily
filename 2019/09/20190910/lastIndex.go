package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"

	fmt.Println(strings.LastIndex(str, "l"))
	fmt.Println(strings.LastIndex(str, "lo"))
	fmt.Println(strings.LastIndex(str, "l0"))
}
