package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("end")
	res, err := http.Get("http://www.google.com/roboys.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	a := "start"
	defer fmt.Println(a)
	a = "end"

}
