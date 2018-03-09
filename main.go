package main

import (
	"log"
)

type Person struct {
	Name string
	Age int
}

func Test(any interface{}) {
	p, ok := any.(Person) 
	if ok {
		log.Printf("casting works %+v", p)
	}
}

func main() {
	p := Person{
		Name: "JP",
		Age: 37,
	}
	Test(p)
}