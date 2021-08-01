package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
	defer wg.Done()
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(s)
		if i == 2 {
			panic("Oh DEAR i==2")
		}
	}

}
func main() {
	wg.Add(2)
	go say("Hey")
	go say("Hello")
	wg.Wait()
}
