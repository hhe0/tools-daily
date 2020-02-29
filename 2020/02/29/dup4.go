package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dupFileNames := make([]string, 0)
	files := os.Args[1:]

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		counts := make(map[string]int)
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if counts[line] > 0 {
				dupFileNames = append(dupFileNames, file)
				break
			}
			counts[line]++
		}
	}

	fmt.Println(len(dupFileNames))
	for _, fileName := range dupFileNames {
		fmt.Println(fileName)
	}
}
