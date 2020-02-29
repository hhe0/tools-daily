package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var s, sep string

	startFor := time.Now().UnixNano()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	endFor := time.Now().UnixNano()
	fmt.Println("使用遍历的方式共需要" + strconv.Itoa(int(endFor-startFor)/1000) + "秒")

	startJoin := time.Now().UnixNano()
	s = strings.Join(os.Args[1:], " ")
	endJoin := time.Now().UnixNano()
	fmt.Println("使用Join的方式共需要" + strconv.Itoa(int(endJoin-startJoin)/1000) + "秒")
}
