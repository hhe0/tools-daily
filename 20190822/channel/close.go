package main

import "fmt"

func callerA2(ch chan string) {
	ch <- "Hello World"
	close(ch)
}

func callerB2(ch chan string) {
	ch <- "你好"
	close(ch)
}

func main() {
	chA, chB := make(chan string), make(chan string)
	go callerA2(chA)
	go callerB2(chB)
	msg := ""
	okA, okB := true, true
	for okA || okB {
		select {
		case msg, okA = <-chA:
			if okA {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, okB = <-chB:
			if okB {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}

