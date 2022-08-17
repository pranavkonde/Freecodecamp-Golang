package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Threads: %v", runtime.GOMAXPROCS(-1))
}

// Shows the number of thread of available

// and if positive then assign that many threads to it
