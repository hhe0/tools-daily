package main

import (
	"fmt"
	"time"
)

func main() {
	//printLetter()
	//printNumbers()
	go printLetter()
	go printNumbers()
	fmt.Println()
}

func printNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d", i)
	}
}

func printLetter() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%c", i)
	}
}
