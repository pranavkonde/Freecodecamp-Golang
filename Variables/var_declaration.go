package main

import "fmt"

var l int = 42

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	docterNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {
	var i int
	i = 42
	i = 27
	fmt.Printf("%v,%T", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v,%T", j, j)

	k := 20

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Println(j, j)
	fmt.Printf("%v,%T", j, j)
}
