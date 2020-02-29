package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int
var ch chan struct{}

func init() {
	ch = make(chan struct{}, 1)
}

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", countHandler)

	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatal(err)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	ch <- struct{}{}
	count++
	fmt.Fprintf(w, "%q\n", r.URL.Path)
	<-ch
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	ch <- struct{}{}
	fmt.Fprintf(w, "%d\n", count)
	<-ch
}
