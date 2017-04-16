package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go hello()
	go world()
	wg.Wait()
}

func hello() {
	for i := 0; i < 45; i++ {
		fmt.Println("hello: ", i)
	}
	wg.Done()
}

func world() {
	for i := 0; i < 45; i++ {
		fmt.Println("world: ", i)
	}
	wg.Done()
}
