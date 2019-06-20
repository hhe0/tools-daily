package main

import "fmt"

type Animal interface {
	hello()
}

type Elephant struct {
	height float32
}

func (elephant Elephant) hello() {
	fmt.Println("I am an elephant")
}

type Dog struct {
	height float32
}

func (dog Dog) hello() {
	fmt.Println("I am a dog")
}

func main() {
	dog := Dog{1}
	elephant := Elephant{3}
	zoo := []Animal{dog, elephant}

	for n := range zoo {

		fmt.Println(zoo[n])
		switch zoo[n].(type) {
		case Elephant:
			fmt.Println("elephant")
		case Dog:
			fmt.Println("dog")
		default:
			fmt.Println("unknown type")
		}
	}
}