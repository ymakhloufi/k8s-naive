package main

import (
	"strconv"
	"time"
)

func main() {
	println("Hello World!")
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		println("I'm alive for " + strconv.Itoa(i) + " seconds...")
	}
	println("Good-Bye World!")
}
