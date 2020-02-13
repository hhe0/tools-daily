package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := "[1]"

	var data []int
	err := json.Unmarshal([]byte(jsonData), &data)
	fmt.Println(err)
	fmt.Println(jsonData, data)
}
