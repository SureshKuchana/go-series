package main

import (
	"fmt"
	"os"
)

// global variables
var username = os.Getenv("USER")

func main(){
	println("Welcome to GO", username);

	// int
	var num int = 10
	var num1 int
	// string
	var str string = "SURESH"
	var str1 string


	// const variables cannot be re-initalized
	// const value = "value"

	fmt.Printf(" str is %s\n", str)
	fmt.Printf(" str1 is %s\n", str1)

	fmt.Printf(" num is %d\n", num)
	fmt.Printf(" num is %d\n", num1)


	sumOfTwoNumbers := addTwo(10 , 20)
	println("Sum of two number is ", sumOfTwoNumbers);

	sum, diff := squareOfNumbers(10, 20)
	println("Square of numbers is ", sum, " ", diff);

	// use the func from another file
	// printData()

	// for loops

	for i := 0; i < 10; i++ {
		println(i)
	}

	// for loop using range
	// for _,value := range nums {
	// 	println(" ll ", value)
	// }

}

func addTwo(x int, y int) int {
	return x + y
}
// or
// func addTwo(x , y int) int {
// 	return x + y
// }

func squareOfNumbers(x, y int) (int , int){
	sum := x + y
	diff := x -y
	return sum, diff
}

// 
