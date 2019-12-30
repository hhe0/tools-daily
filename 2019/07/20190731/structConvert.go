package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int
}

type Student struct {
	Name  string
	Age   int
	StdNo int
}

func main() {
	student := Student{
		Name:  "hehong",
		Age:   24,
		StdNo: 123456789,
	}

	fmt.Println(student)
	// 转换错误
	//Person(student)

	jsonData, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}

	var person Person
	err = json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person)
}
