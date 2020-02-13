package main

import "fmt"

func main() {
	strA := `hello`
	fmt.Println(strA[1:3])
	fmt.Println(len(strA))

	strB := "\xe4\xb8\x96"
	fmt.Println(strB)
	fmt.Println(len(strB))

	strC := "\u4e16"
	fmt.Println(strC)

	strD := "世"
	fmt.Printf("% x\n", strD)
	fmt.Printf("% x\n", []rune(strD))

	fmt.Printf("%x \n", []byte(strD))

	strDArr := []byte(strD)
	for _, v := range strDArr {
		fmt.Println(v)
	}
}
