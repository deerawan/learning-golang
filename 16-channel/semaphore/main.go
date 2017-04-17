package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true
	}()

	go func() {
		for i := 100; i < 110; i++ {
			ch <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(ch)
	}()

	// This code is blocking so, it will wait until c finished
	for n := range ch {
		fmt.Println(n)
	}
}
