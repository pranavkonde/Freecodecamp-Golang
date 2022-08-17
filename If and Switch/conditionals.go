package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}
	if false {
		fmt.Println("The test is true")
	}
	statePopulations := map[string]int{
		"California":  34567,
		"Texas":       345667,
		"Florida":     987654,
		"New York":    9876543,
		"Pennsyvania": 45678,
		"Ohio":        876543,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	// comparison operators works here maily

	number := 50
	guess := -5
	if guess < 1 || returnTrue() || number > 100 { // will only check first case as in or one condition must be true this concept is called short curcuiting
		fmt.Println("The guess must be between 1 and 100")
	}

	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}

func returnTrue() bool {
	fmt.Println("Inside returnTrue function")
	return true
}

/* if condition{

} else if condition {

} else {

}
*/
