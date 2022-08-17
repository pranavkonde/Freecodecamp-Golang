package main

import "fmt"

// In map ordering get changes
func main() {
	statePopulations := map[string]int{
		"California":  34567,
		"Texas":       345667,
		"Florida":     987654,
		"New York":    9876543,
		"Pennsyvania": 45678,
		"Ohio":        876543,
	}
	fmt.Println(statePopulations)
	m := map[[3]int]string{}
	fmt.Println(statePopulations, m)

	statePopulations1 := make(map[string]int)
	statePopulations1 = map[string]int{
		"California":  34567,
		"Texas":       345667,
		"Florida":     987654,
		"New York":    9876543,
		"Pennsyvania": 45678,
		"Ohio":        876543,
	}
	fmt.Println(statePopulations1)
	fmt.Println(statePopulations1["Ohio"])
	statePopulations1["Georgia"] = 10310371
	delete(statePopulations1, "Georgia")
	fmt.Println(statePopulations1)
	fmt.Println(statePopulations1["Georgia"]) // if not present map will never show error it simply prints zero
	pop, ok := statePopulations1["Oho"]       //will check wheather Oho is present in map or not and return true or false accordingly
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations1))

	sp := statePopulations1
	delete(sp, "Ohio")
	fmt.Println(statePopulations1)

}
