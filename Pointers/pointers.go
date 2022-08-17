package main

import "fmt"

type myStruct struct {
	foo int
}
type hiStruct struct {
	loo int
}

func main() {
	a := 42
	b := a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
	var c *int = &a
	fmt.Println(c)
	fmt.Println(*c)

	p := [3]int{1, 2, 3}
	q := &p[0]
	r := &p[1]
	fmt.Printf("%v %p %p", p, q, r)

	var ms myStruct
	ms = myStruct{foo: 42} //new(myStruct) intializes the values as 0
	fmt.Print(ms)

	var mn *hiStruct
	mn = new(hiStruct)
	(*mn).loo = 43         //mn.loo
	fmt.Println((*mn).loo) //mn.loo

}
