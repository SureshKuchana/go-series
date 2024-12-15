package main

import "fmt"

// global variable can declared outside of the func
var url = "www.google.com"

func main(){
	var x int	
	var name string
	var isBool bool
	const y = 2
	var x1 int = 2
	var name1 string = "Suresh"
	var isBool1 bool = true
	x = 10
	name1 = "Suresh K"
	// variable initialization shortcut
	otherText := "Bye!"
		

	print(" Hello from module")
	print(x, " x before")
	print(name, " name ")
	print(isBool, " isBool ")
	print(y, " y ")
	print(x1, " x1 ")
	print(name1, " name1 ")
	print(isBool1, " isBool1 ")
	print(x, " x after")
	print(otherText, " otherText")
	print(" google ", url)
	var fn string
	print(" what is your name: " )
	fmt.Scanf("%s ", &fn )
	print(" Hello ", fn )

}