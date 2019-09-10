package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "*hello world****"
	fmt.Println(strings.Trim(str, "*"))
	fmt.Println(strings.Trim(str, "**"))
	fmt.Println(strings.Trim(str, "*h"))
}
