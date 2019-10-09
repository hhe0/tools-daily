package main

import "fmt"

func main() {
	//sn1 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//sn2 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//fmt.Println(sn1 == sn2)
	//
	//sn3 := struct {
	//	age  int
	//	name string
	//}{age: 11}
	//sn4 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: ""}
	//fmt.Println(sn3 == sn4)

	//sn5 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//sn6 := struct {
	//	name string
	//	age  int
	//}{age: 11, name: "qq"}
	//fmt.Println(sn5 == sn6)

	sm7 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm8 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	fmt.Println(sm7 == sm8)
}
