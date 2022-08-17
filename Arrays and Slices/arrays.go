// 1:48:01

package main

import "fmt"

func main() {
	grade1 := 97
	grade2 := 85
	grade3 := 93
	fmt.Printf("Grades: %v,%v,%v\n", grade1, grade2, grade3)

	grades := [3]int{95, 96, 97}
	fmt.Printf("Grades: %v\n", grades)

	grades1 := [...]int{95, 96, 97, 98, 99, 100} // we can also keep empty which will work same as ...
	fmt.Printf("Grades: %v\n", grades1)

	var student [3]string
	fmt.Printf("Students: %v\n", student)
	student[0] = "Hello"
	student[1] = "Hello1"
	student[2] = "Hello2"
	fmt.Printf("Students: %v\n", student)
	fmt.Printf("Student #1: %v\n", student[1])
	fmt.Printf("Number of Students:%v\n", len(student))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	var identityMatrix1 [3][3]int
	identityMatrix1[0] = [3]int{1, 0, 0}
	identityMatrix1[1] = [3]int{0, 1, 0}
	identityMatrix1[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix1)

	a := [...]int{1, 2, 3}
	b := a // Array are passed by value (i.e if value of b is changed it will not affect in a)
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)

}
