package main

import "fmt"

func main() {
	mapA := make(map[string]string)
	mapB := make(map[string]string)
	fmt.Print(mapA == mapB)
}
