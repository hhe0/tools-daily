package main

import (
	"encoding/json"
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
		"bag": map[string]interface{}{
			"size":  20,
			"color": "red",
			"has":   true,
		},
	}

	fmt.Println(people)

	peopleTrans := MapTransformer(people)
	fmt.Println(peopleTrans)

	peopleJson, _ := json.Marshal(peopleTrans)
	fmt.Println(string(peopleJson))
}

func MapTransformer(params map[string]interface{}) map[string]interface{} {
	keys := make([]string, 0)
	for element, value := range params {
		if v, ok := value.(map[string]interface{}); ok {
			// 如果还是一个map，递归处理
			params[element] = MapTransformer(v)
		}
		keys = append(keys, element)
	}

	sort.Strings(keys)

	newParams := make(map[string]interface{}, 0)
	for _, key := range keys {
		if element, ok := params[key]; ok {
			newParams[key] = element
		}
	}

	return newParams
}
