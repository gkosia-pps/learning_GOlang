package main

import "fmt"

func main_03() {

	var my_int_arry = [5]int{}

	var my_string_array [3]string

	my_int_arry[0] = 99

	my_string_array[0] = "zero"

	fmt.Println(len(my_int_arry))
	fmt.Println(my_int_arry)
	fmt.Println(my_string_array[0])

	var my_slice []string

	my_slice = append(my_slice, "Gab")

	fmt.Println(my_slice)
	fmt.Printf("%T", my_slice)
}
