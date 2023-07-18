package main

import (
	"fmt"
	"strings"
)

func main_04() {

	names := []string{"A a", "B b", "C c"}

	// infinite loops
	// for { <code>}

	// for each item of array or slice
	for _, val := range names {
		fmt.Println(strings.Fields(val))
	}

}
