package main

import "fmt"

// In go there is no need to add break in each case
func main() {
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2, 5:
		fmt.Println("two")
	case 3:
		fmt.Println("not one or two")
	}
	/* switch i:=5; i{

	}
	*/
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than equal to ten")
	case i <= 20:
		fmt.Println("Less than equal to twenty")
	default:
		fmt.Println("greator than twenty")
	}
	// fallthrough keyword is used to execute the down case below the fallthrough case, and does not matter wheather the down case is true or false
	switch {
	case i <= 10:
		fmt.Println("less than equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("Less than equal to twenty")
		fallthrough
	case i <= 30:
		fmt.Println("Less than equal to thirty")
	default:
		fmt.Println("greator than twenty")
	}

	var z interface{} = 1
	switch z.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}
}
