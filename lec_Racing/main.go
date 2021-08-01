package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Cpu Num", runtime.NumCPU())
	fmt.Println("Goroutine Num", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("Cpu Num", runtime.NumCPU())
	fmt.Println("Goroutine Num", runtime.NumGoroutine())
}
