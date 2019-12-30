package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	testRuntime()
	testTime()
	testBlank()
}

func testRuntime() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

func testTime() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	fmt.Printf("%T\n", today)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func testBlank() {
	t := time.Now()
	// 等价于switch true
	// 这一构造使得可以用更清晰的形式来编写长的if-then-else链
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}
}

