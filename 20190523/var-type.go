package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	i int = 64
	// 类型转换必须做显示转换，否则会出错
	a float64 = float64(i)
)

const (
	PI = 3.14
	Big = 1 << 100
	Small = Big >> 99
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe)
	fmt.Printf(f, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Printf(f, a)

	v := 42.12
	// int float64 complex128
	fmt.Printf("v is of type %T\n", v)

	fmt.Println(PI);

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}