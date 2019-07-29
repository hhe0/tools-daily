package main

import "fmt"

type Student struct {
	Age  int
	Name string
}

func main() {
	student := Student{
		Age: 23,
		Name: "hehong",
	}

	fmt.Println((&student).Age)
	fmt.Println(*(&(student.Age)))
}

