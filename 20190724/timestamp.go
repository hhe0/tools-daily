package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// 当前时间 Time类型
	fmt.Println(reflect.TypeOf(time.Now()), time.Now())

	// 当前时间戳 秒级
	fmt.Println(reflect.TypeOf(time.Now().Unix()), time.Now().Unix())

	// 当前时间戳 纳秒级
	fmt.Println(reflect.TypeOf(time.Now().UnixNano()), time.Now().UnixNano())

	// 时间戳转日期
	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))

	// 当前年、月、日、...
	fmt.Println(reflect.TypeOf(time.Now().Year()), time.Now().Year())

	// 时间戳运算
	fmt.Println(time.Now().Add(1 * time.Minute).Unix())
}



