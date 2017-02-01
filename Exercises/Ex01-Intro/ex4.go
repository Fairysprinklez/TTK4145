package main

import (
	. "fmt"
	"time"
)

var Global = 0

func routine1() {
	for i := 0; i < 1000000; i++ {
		Global++
	}

}

func routine2() {
	for i := 0; i < 1000000; i++ {
	Global--
	}
}

func main() {
	go routine1()
	go routine2()

	time.Sleep(100*time.Millisecond)
	Println(Global)
}
