package main

import (
	"fmt"
	"time"
)

func say(s string, done chan string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	done <- "Terminei"
}

func main() {
	done := make(chan string, 1)

	//go say("world", done)
	done <- "Teste"
	//
	fmt.Println(<-done)
}
