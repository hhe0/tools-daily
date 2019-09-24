package main

import (
	"encoding/json"
	"fmt"
)

type jsonStruct struct {
	IntData    int        `json:"int_data"`
	StringData string     `json:"string_data"`
	arrData    []int
	StructData structData `json:"struct_data"`
}

type structData struct {
	Data int
}

func main() {
	jsonStr := `{"int_data1":10,"STRING_DATA":"hehong","structdata":{"Data":3}}`
	var res jsonStruct
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	a := `{"key": 111}`
	var b int
	_ = json.Unmarshal([]byte(a), &b)
	fmt.Println(b)
}
