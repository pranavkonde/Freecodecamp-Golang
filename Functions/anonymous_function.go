package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello Go!")
	}()
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
	f := func() {
		fmt.Println("Hello Go!")
	}
	f()
	var q func() = func() {
		fmt.Println("Hello Go!")
	}
	q()
}
