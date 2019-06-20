package main

import "fmt"

type SimpleInterface interface {
	Get() int
	Set(int)
}

type SimpleStruct struct {
	v int
}

func (s *SimpleStruct) Get() int {
	return s.v
}

func (s *SimpleStruct) Set(a int) {
	s.v = a * a
}

func main() {
	var aInf SimpleInterface
	a := &SimpleStruct{3}
	aInf = a
	c := aInf.Get()
	fmt.Println(c)

}