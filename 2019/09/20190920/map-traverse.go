package main

import "fmt"

func main() {
	people := map[string]interface{}{
		"name":   "hehong",
		"sex":    "男",
		"age":    24,
		"height": 177,
		"weight": 77,
	}

	fmt.Println(people)

	for k, v := range people {
		fmt.Println(k, v)
	}
}
