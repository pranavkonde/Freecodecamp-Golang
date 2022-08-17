package main

import "fmt"

const a int16 = 32

// we can change value of const when redeclare in function along with the data type
const (
	e = iota /* can be replaced to e=iota f g  but value of e=0,f=1,g=2*/
	f = iota // if we add e+5 then f and g would be 6 and 7
	g = iota
)

const (
	catSpecialist = iota //iota data type is int
	dogSpecialist
	snakeSpecialist
)
const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	const myConst int = 42 // Must be assign value at compile time
	fmt.Printf("%v,%T", myConst, myConst)

	const a int = 14 // const a = 42
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fileSize := 4000000000
	fmt.Printf("%.2fGB", fileSize/GB)
}

// 1:45:28
