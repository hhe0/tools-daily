package main

import (
	"fmt"
	"sort"
)

func main() {
	people := map[string]interface{}{
		"name":   "hehong",
		"sex":    "男",
		"age":    24,
		"height": 177,
		"weight": 77,
	}

	keys := make([]string, 0)
	for v := range people {
		keys = append(keys, v)
	}
	fmt.Println(keys)

	sort.Strings(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, people[key])
	}
}
