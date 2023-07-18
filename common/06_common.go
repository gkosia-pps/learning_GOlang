package common

import "fmt"

// To make a functnio public we need to export it
// To export it we change the first letter to capital and becomes public
func Print_something_common() {
	fmt.Println("This is common to all")
}
