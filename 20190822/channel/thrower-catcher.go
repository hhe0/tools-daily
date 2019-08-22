package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw <<", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <- c
		fmt.Println("Catcher >>", num)
	}
}

func main() {
	ch := make(chan int, 2)
	go thrower(ch)
	go catcher(ch)
	time.Sleep(10 * time.Second)
}