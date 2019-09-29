package main

import "fmt"

func main() {
	d := Dog{
		Pet{},
	}
	d.Skill()
	d.Navigate()

	c := Cat{
		Pet{"leo"},
	}
	c.Skill()
	c.Catch()
	fmt.Println(c.Name)
}

type Pet struct {
	Name string
}

func (p Pet) Skill() {
	fmt.Println("能文能武的宠物")
}

type Cat struct {
	Pet
}

func (c Cat) Catch() {
	fmt.Println("老鼠天敌喵喵喵")
}

type Dog struct {
	Pet
}

func (d Dog) Navigate() {
	fmt.Println("自带导航汪汪汪")
}

func (d Dog) Skill() {
	fmt.Println("覆写方法")
}
