package main

import "fmt"

func main() {

	// we will manipulate some basic functions and features of go lang

	var somethingNew string
	var i int
	// variable declaration
	somethingNew = "This is a new learning"

	i = 4

	fmt.Println(somethingNew)

	fmt.Println("The value that I want to print is", i)

	// the output can be stored in different variables seperated by comma
	// how to use both values at different places, ... we will learn in future
	// := is a new way of declare variable and assign values
	someNewFunctionString, anotherVariable := aNewFunction()

	fmt.Println("This lines are output of another function with 2 return values to print", someNewFunctionString, anotherVariable)
}

// it is returning string but it can return more than 1 values by sepearting the types by comma like string, string
//and putting inside (). for 1 return value () is not required
func aNewFunction() (string, string) {
	return "something", "else"
	//function is returning 2 strings

}
