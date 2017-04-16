package main

import "fmt"
import "sync"
import "runtime"

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

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
