package main

import "fmt"

var my_package_level_var = "Package var"

// return multiple values (int, int ,int)
func print_something(name string) (int, int, int) {
	fmt.Println("Hey", name, "This is a msg from print_something func", my_package_level_var)

	return 1, 2, 3
}
func main() {

	fmt.Println(print_something("Gab"))

}
