// just a file of go scraps for copy paste into later programs

package main

// importing libraries
import "fmt"

// Variables
var name = "Brad"
var age int = 19
var isCool = true
var isnotCool = false
var name = "Jason"
name := "Jason"  // you can only use this inside of a function
var email = "jason@gmail.com"
name,email := "jason","jason@gmail.com"

// run the main function
func main() {

	// print the string
	fmt.Println("Starting Program...")
	// get the variable type of name
	fmt.Printf("%T\n", name)
	//	print_name()
}

// another function
func print_name() {
	fmt.Println(name, age)
}

//func main() {
//	fmt.Println("Hello World")
//}
