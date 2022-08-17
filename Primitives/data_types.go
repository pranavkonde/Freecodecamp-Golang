package main

import (
	"fmt"
)

// By default numeric values are zero and boolean are false

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	a := 1 == 1
	b := 2 == 3
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	var c bool
	fmt.Println(c)

	// unsigned int are always positive and signed int can be positive or negative

	var d uint16 = 42
	fmt.Printf("%v, %T", d, d)

	// int + int8 = ERROR type conversion is to be done
	/* a:=8
	a<<3  (2^3*2^3) = 2^6 =64
	a>>3 2^3/2^3=1
	*/

	e := 3.14
	e = 13.7e72
	e = 2.1e14
	fmt.Printf("%v ,%T", e, e)

	f := 10.2
	g := 3.7
	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)

	// complex are only 64 and 128
	var h complex64 = 1 + 2i
	fmt.Printf("%v,%T", h, h)

	var j complex64 = 1 + 2i
	fmt.Printf("%v,%T", real(j), real(j))
	fmt.Printf("%v,%T", imag(j), imag(j))

	var k complex64 = complex(5, 12)
	fmt.Printf("%v,%T", k, k)

	s := "Hello"
	fmt.Printf("%v,%T", s, s)

	s = "Hello World"
	fmt.Printf("%v,%T", s[1], s) // will return alias of bytes
	fmt.Printf("%v,%T", string(s[1]), s)

	t := "Josh"
	fmt.Printf("%v,%T\n", s+t, s+t)

	u := []byte(s) // will make a slice ASCII values from string
	fmt.Printf("%v,%T\n", u, u)

	y := 'a'
	fmt.Printf("%v,%T", y, y)

}
