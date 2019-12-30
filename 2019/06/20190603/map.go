package main

type Student struct {
	age int
	sex string
}

var a map[string]Student

func main() {
	// 放在这里会有问题？？
	//var a map[string]Student
	a["0"] = Student{23, "female"}


	m := make(map[string]Student)
	m["111"] = Student{23, "male"}

}
