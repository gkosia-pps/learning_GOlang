package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func main_08() {

	var gab = Person{
		name: "Gab",
		age:  30,
	}

	fmt.Println(gab)
	fmt.Println(gab.name)
}
