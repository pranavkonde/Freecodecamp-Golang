package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg3 = sync.WaitGroup{}
var counter3 = 0
var m = sync.RWMutex{} // one can read but many can write at a time

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg3.Add(2)
		go sayHello3()
		go increment3()
	}
	wg3.Wait()
}
func sayHello3() {
	m.RLock()
	fmt.Printf("Hello %v\n", counter3)
	m.RUnlock()
	wg3.Done()
}
func increment3() {
	m.Lock()
	counter3++
	m.Unlock()
	wg3.Done()
}
