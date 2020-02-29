package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		fmt.Println(resp.StatusCode)
	}
}
