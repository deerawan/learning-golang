package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // close the channel, so nobody can write on it again
	}()

	// This code is blocking so, it will wait until c finished
	for n := range ch {
		fmt.Println(n)
	}
}
