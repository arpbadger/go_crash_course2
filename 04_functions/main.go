// main go is a template for learning functions in golang
package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	//fmt.Println("Welcome")
	//fmt.Println(greeting("Brad"))
	fmt.Println(getSum(3, 4))
}
