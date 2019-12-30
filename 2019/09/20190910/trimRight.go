package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "*hello world****"
	fmt.Println(strings.TrimRight(str, "*"))
	fmt.Println(strings.TrimRight(str, "**"))
	fmt.Println(strings.TrimRight(str, "*h"))
}
