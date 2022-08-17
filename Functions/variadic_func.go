package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 100)
}
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is: ", result)
}
