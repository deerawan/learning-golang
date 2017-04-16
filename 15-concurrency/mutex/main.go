package main

import "fmt"
import "sync"
import "time"

// using go run -race main.go => to detect race condition

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
	wg.Add(2)
	go incrementor("foo")
	go incrementor("bar")
	wg.Wait()

	fmt.Println("Final Counter", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(5 * time.Millisecond))
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter", counter)
		mutex.Unlock()
	}
	wg.Done()
}
