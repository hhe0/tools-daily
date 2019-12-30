package main

import (
	"fmt"
	"strings"
)

func main() {
	strSlice := []string{"hello", "world"}
	fmt.Println(strings.Join(strSlice, " "))
}
