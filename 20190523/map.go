package main

import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

func main() {
	// map在使用之前必须使用make而不是new来创建
	// 值为nil的map是空的，并且不能赋值
	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.68433,
		-74.39967,
	}
	fmt.Println(m["Bell Labs"])

	delete(m, "Bell Labs")
	m["Leo Labs"] = Vertex1{
		32.111,
		-21.21212,
	}
	fmt.Println(m)
}
