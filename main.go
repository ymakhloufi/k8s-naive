package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func printstdout() {
	println("Hello World!")
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		println("I'm alive for " + strconv.Itoa(i) + " seconds...")
	}
	println("Good-Bye World!")
}

func main() {
	go printstdout()

	http.HandleFunc("/", healthy)
	http.HandleFunc("/health", healthy)
	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":8090", nil)
}

func healthy(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "healthy...")
}

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong - "+time.Now().Local().String())
}
