package main

import (
	"fmt"
	"time"
)

func callerA(ch chan string) {
	ch <- "Hello World"
}

func callerB(ch chan string) {
	ch <- "你好"
}

func main() {
	chA, chB := make(chan string), make(chan string)
	go callerA(chA)
	go callerB(chB)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)
		select {
		case msg := <-chA:
			fmt.Printf("%s from A\n", msg)
		case msg := <-chB:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("Default")
		}
	}
}
