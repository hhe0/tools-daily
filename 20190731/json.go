package main

import (
	"encoding/json"
	"fmt"
)

type jsonStruct struct {
	IntData    int        `json:"int_data,omitempty"`
	StringData string     `json:"string_data"`
	arrData    []int      `json:"arr_data"`
	StructData structData `json:"struct_data"`
}

type structData struct {
	Data int
}

func main() {
	structData := jsonStruct{
		IntData:    0,
		StringData: "hehong",
		arrData:    []int{1, 2, 3},
		StructData: structData{Data: 3},
	}

	byteArrData, _ := json.Marshal(structData)
	jsonData := string(byteArrData)
	fmt.Println(string(jsonData))

	jsonStr := "{\"string_data\":\"hehong\",\"struct_data\":{\"Data\":3}}"
	var res jsonStruct
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		fmt.Println(err)
	}
}
