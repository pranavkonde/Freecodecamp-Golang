package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	q := 0
	for ; q < 5; q++ {
		fmt.Println(q)
	}
	for q < 5 {
		fmt.Println(q)
	}
	for q < 5 {
		fmt.Println(q)
	}
Loop:
	for l := 1; l <= 3; l++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(l * j)
			if l*j >= 3 { // this is break the outer for loop
				break Loop
			}
		}
	}
	s := []int{1, 2, 3}
	fmt.Println(s)

	for u, v := range s {
		fmt.Println(u, v)
	}

	t := "Hello Go!"
	for u, v := range t {
		fmt.Println(u, string(v))
	}

}
