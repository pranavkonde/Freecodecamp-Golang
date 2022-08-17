package main

import "fmt"

func main() {

	c := []int{1, 2, 3}
	d := c // Slices are pass by reference (i.e if value of d is changed it will affect the value of c)
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c))

	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	q := p[:]   // slice of all elements
	r := p[3:]  //slice from 4th element to end
	s := p[:6]  //slice first 6 elements
	t := p[3:6] //slice the 4th,5th and 6th elements
	p[5] = 42   //Will make change in all the slices
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println(t)

	a := make([]int, 3, 100) // Way to make a slice using make (make(type,len,cap))
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	x := []int{}
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))

	x = append(x, 1)
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))

	x = append(x, 2, 3, 4, 5) //capcity doubles as we add element into slice as 1-2-4-8-16
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))

	x = append(x, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}...)
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))

	y := []int{1, 2, 3, 4, 5}
	z := y[1:]
	fmt.Println(z)

	l := []int{1, 2, 3, 4, 5}
	m := l[:len(y)-1]
	fmt.Println(m)

	u := []int{1, 2, 3, 4, 5}
	v := append(u[:2], u[3:]...)
	fmt.Println(u) // value of u gets change when appending of v is done
	fmt.Println(v)
}

//2:12:40
