package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := 1564559040

	fmt.Println(transformerDate(timestamp))
}

func transformerDate(timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format("2016-01-02 15:04:05")
}
