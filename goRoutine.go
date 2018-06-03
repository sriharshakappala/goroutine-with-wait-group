package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Printf("In Go Routine %d \n", i)
}

func changeToMain() {
	for i := 0; i < 1000; i++ {
		go hello(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("In main")
}
