package main

import "fmt"
import "sync"
import "time"

var wg sync.WaitGroup
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
		counter++
		time.Sleep(time.Duration(5 * time.Millisecond))
		fmt.Println(s, i, "Counter", counter)
	}
	wg.Done()
}
