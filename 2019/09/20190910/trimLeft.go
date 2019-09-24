package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "*hello world****"
	fmt.Println(strings.TrimLeft(str, "*"))
	fmt.Println(strings.TrimLeft(str, "**"))
	fmt.Println(strings.TrimLeft(str, "*h"))
}
