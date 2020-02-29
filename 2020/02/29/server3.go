package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	count2 int
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/", handler3)
	http.HandleFunc("/count", handlerCount2)

	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatal(err)
	}
}

func handler3(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count2++
}

func handlerCount2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(w, "%d\n", count2)
}
