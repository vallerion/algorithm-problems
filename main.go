package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	var done bool

	go func() {
		for !done {
		}
		fmt.Println(x)
	}()

	go func() {
		x = 1
		done = true
	}()

	time.Sleep(time.Millisecond)
}
