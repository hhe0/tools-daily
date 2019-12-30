package main

import "fmt"

type Iname interface {
	Mname()
}
type St struct{}

func (St) Mname() {}
func main() {
	var t *St
	if t == nil {
		fmt.Println("t is nil")
	} else {
		fmt.Println("t is not nil")
	}
	var i Iname = t
	fmt.Printf("%T\n", i)
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is not nil")
	}
	fmt.Printf("i is nil pointer:%v", i == (*St)(nil))
}
