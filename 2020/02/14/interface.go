package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) String() {
	fmt.Println(s.Name + strconv.Itoa(s.Age))
}

func main() {
	s := Student{"hehong", 25}
	fmt.Println(fmt.Print(s))
}
