package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"` //tags it wll suggest that the value of name is mandatory and the max size of name should be less than or equal to 100
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
