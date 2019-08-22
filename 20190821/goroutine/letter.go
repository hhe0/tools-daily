package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers()
	printLetter()
}

func printNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func printLetter() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c", i)
	}
}
