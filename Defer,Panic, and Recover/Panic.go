package main

import "fmt"

func main() {
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")

	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}
