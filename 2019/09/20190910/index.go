package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	fmt.Println(strings.Index(str, "l"))
	fmt.Println(strings.Index(str, "lo"))
	fmt.Println(strings.Index(str, "l0"))
}
