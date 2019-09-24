package main

import "fmt"

func main() {
	fmt.Println(Multi(1, 3))
}

func Multi(x, y int) (sum int, err error) {
	return x + y, nil
}
