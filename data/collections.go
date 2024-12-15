package data

import "fmt"

// array have fixed length
var Countries [10]string
var arr [1]string

// slices are dynamic length
var slices []string

// map is key value pair
var maps map[string]int

func init(){
	fmt.Print(Countries)
	Countries[0] = "A"
	Countries[1] = "B"
	Countries[2] = "C"

	slices[0] = "Hello"
	slices[1] = "World"

	names := append(slices, "Great")

	fmt.Print(Countries)
	fmt.Print(names)
	fmt.Print(arr)
}