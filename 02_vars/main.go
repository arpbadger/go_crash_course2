package main

import "fmt"

// create variable using var keyword
//var name string = "Brad"
var name = "Brad"
var age int = 19

// run the main function
func main() {
	//short hand
	name := "Jason"

	// print the string
	fmt.Println("Starting Program...")
	fmt.Println(name)
	// get the variable type of name
	//fmt.Printf("%T\n", name)
	//	print_name()
}

// another function
func print_name() {
	fmt.Println(name, age)
}

//func main() {
//	fmt.Println("Hello World")
//}
