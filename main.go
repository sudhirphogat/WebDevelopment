package main

import (
	"log"
	"time"
)

//we can define our own type using "type"
//we  will use struct to group variables

//example
//below list of variables will increase the list and hectic to remember and maintain
//******the first letter is CAPS because we can make the variable global to other packages by it
//********* first small letter variable scope is inside the package only
//var FirstName string
//var LastName string
//var age int
//var PhoneNumber string
//var DateOfBirth time.Time

//***use struct to define variable
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	DateOfBirth time.Time
}

//Below variables we will see conversion
var a int

type checkConvertion int

var b checkConvertion

func main() {
	//struct variable can be difined like this
	//values that are not used are assigned default values
	//********the variable user is defined like "**var user User**"

	user := User{
		FirstName: "Sudhir",
		LastName:  "Phogat",
		//end it with a comma
	}

	log.Println(user.FirstName, user.LastName, "Birthday=", user.DateOfBirth)

	a = 3
	b = 4
	log.Println(a)
	log.Printf("%T\n", a)
	log.Println(b)
	log.Printf("%T\n", b)

	// b cannot be assigned to a as a is int and b is checkConversion
	//a=b

	// ***but it can be converted to another type as the underline of b is int
	a = int(b)
	log.Println(a)
	log.Printf("%T\n", a)

	//**** multiple ways of defining a structure

	var myVar User
	myVar.FirstName = "Sudhir"

	myVar2 := User{
		FirstName: "Phogat",
	}

	//log.Println("The value of myVar is", myVar.FirstName)
	//log.Println("The value of myVar2 is", myVar2.FirstName)

	//****using function instead of direct variable
	log.Println("The value of myVar is", myVar.printFirstName())
	log.Println("The value of myVar2 is", myVar2.printFirstName())
}

//***Type Struct can also be called using a function by passing  "* struct" before the func name
//****KeyNote** before function (m *User) where User is type and struct  and after function (s *string) means pointer

func (m *User) printFirstName() string {
	return m.FirstName
}
