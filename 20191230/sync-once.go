package main

import (
	"fmt"
	"sync"
)

var times int

func main() {
	var once sync.Once

	once.Do(printHello)
	once.Do(printHello)
	once.Do(printHello)
	once.Do(printHello)
	once.Do(printHello)
}

func printHello() {
	fmt.Println("Hello")

	times++
	fmt.Println(times)
}
