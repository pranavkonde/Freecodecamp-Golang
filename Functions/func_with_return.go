package main

import "fmt"

func main() {
	s := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The sum is: ", s)
}
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
