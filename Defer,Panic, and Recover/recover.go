package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Panicln("Error:", err)
		}
	}()
	panic("Something bad happenend")
	fmt.Println("end")
}
