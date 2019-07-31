package main

import (
	"fmt"
)

func main() {
	name := "Minc"
	chineseName := "何宏"

	fmt.Println(name)
	fmt.Println(chineseName)

	fmt.Println(name[0:1])
	fmt.Println(chineseName[0:1])


	fmt.Println([]byte(name))
	fmt.Println([]byte(chineseName))

	fmt.Println(string([]byte(name)[0:1]))
	fmt.Println(string([]byte(chineseName[0:1])))

	fmt.Println([]rune(name))
	fmt.Println([]rune(chineseName))
	fmt.Println(string([]rune(name)[0:1]))
	fmt.Println(string([]rune(chineseName)[0:1]))
}
