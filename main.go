package main

import (
	"fmt"
	"log"
)

//this variable ouside function' scope is entire package
var s = "seven"

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

	fmt.Println("This line is output of function aNewFunction with 2 return values to print", someNewFunctionString, anotherVariable)

	//in below example we will be learning pointer and it is similar to c programming
	var myPointerString string
	myPointerString = "Green"

	log.Println("The pointer string is set to ", myPointerString)
	//when we send value to the fucntion it will show an error
	//so to send the reference to pointers or memory location to the function we add "&" infront of the string
	changeUsingPointer(&myPointerString)
	//check the value of the string after the function.
	//Note: there is no return in the function but this happens only because the memory value is changed
	log.Println("the new value of myPointerString after the pointer function is", myPointerString)

	//saysomething function

	sayValue, sayValue2 := saySomething("xxx")
	log.Println("the output of the saysomething funtion is", sayValue, sayValue2)

}

// it is returning string but it can return more than 1 values by sepearting the types by comma like string, string
//and putting inside (). for 1 return value () is not required
func aNewFunction() (string, string) {
	return "something", "else"
	//function is returning 2 strings

}

//apply "*" infront of string which means the location
//so it will not get the value of the value but the memory location
func changeUsingPointer(s *string) {
	// the value that s gets is the memory location called pointer
	log.Println("s is set to ", s)
	newValue := "Red"
	//now set the new value in the location by applying * before the string.
	*s = newValue
	// Note: -  That we are not returning anything but just changing the value in the memory location
}

// now the same "s" is a variable in the package scope and we are getting a value of s in function
// passing value to this function in main function
func saySomething(s string) (string, string) {
	//the value of package scope variable is overwritern when passing the value to the function
	//if the variable name in function is changed like "s to s2"
	//than value of s will be the value of variable in package scope
	log.Println("the value of s in saysomething function is", s)
	return s, "world"

}
