package main

import "fmt"
import "sync"

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 100; i < 110; i++ {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	// This code is blocking so, it will wait until c finished
	for n := range ch {
		fmt.Println(n)
	}
}
