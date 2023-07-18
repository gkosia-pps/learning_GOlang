package main

import (
	"fmt"
)

var my_package_level_variable = "Package var"

func main_01() {

	var cust_name string = "Gab"
	cust_city := "Lim"

	const constVar = "CANNOT_CHANGE"

	fmt.Printf("The type of cust_name is %T\n", cust_name)

	fmt.Println("Hello", cust_name, "wellcome to the app,", constVar)

	fmt.Printf("Welcome %v, from city %v, constant %v \n", cust_name, cust_city, constVar)

	// this is a comment
	var user_age int

	fmt.Printf("Cust age is %v", user_age)
}
