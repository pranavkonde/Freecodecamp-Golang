package main

import "fmt"

type Docter struct {
	number    int
	actorName string
	companies []string
}

func main() {
	aDocter := Docter{
		number:    3,
		actorName: "John Pertwee",
		companies: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	bDocter := Docter{
		3, // This way is not used most
		"John Pertwee",
		[]string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDocter)
	fmt.Println(bDocter)
	fmt.Println(aDocter.actorName)
	fmt.Println(aDocter.companies[1])

	cDocter := struct{ name string }{name: "John Pertwee"}
	fmt.Println(cDocter)
	anotherDocter := &cDocter
	anotherDocter.name = "Tom Baker"
	fmt.Println(cDocter)
	fmt.Println(anotherDocter)
}
