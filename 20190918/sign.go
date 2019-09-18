package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	oldMap := map[string]string{"name": "hehong", "age": "1", "size": "18", "height": "180"}
	fmt.Print(oldMap)

	keys := make([]string, 0)
	for element := range oldMap {
		keys = append(keys, element)
	}

	fmt.Print(keys)

	sort.Strings(keys)

	fmt.Println(keys)

	newMap := make(map[string]string, len(oldMap))
	for _, key := range keys {
		newMap[key] = oldMap[key]
	}

	newMapJson, _ := json.Marshal(newMap)
	fmt.Println(string(newMapJson))
}
