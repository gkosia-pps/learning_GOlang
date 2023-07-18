package main

import (
	"fmt"
	"strconv"
)

func main_07() {

	// map[<datatype of key>]<datatype of value>
	var cust_map = make(map[string]string)

	cust_map["name"] = "Gab"
	cust_map["age"] = strconv.FormatInt(30, 10)

	fmt.Println(cust_map)

	// slicer of maps: [] slicer map[string]string map, 0 initial size
	// var list_of_maps = make([]map[string]string,0)
}
