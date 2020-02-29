package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urls := os.Args[1:]
	start := time.Now()

	ch := make(chan string)
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	seconds := time.Since(start).Seconds()
	fmt.Printf("%.7fs elapsed\n", seconds)
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v\n", err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%v\n", err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.7fs %7d %s", seconds, nbytes, url)
}
