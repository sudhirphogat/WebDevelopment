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

func main() {
	//struct variable can be difined like this
	//values that are not used are assigned default values
	//********the variable user is defined like "**var user User**"

	user := User{
		FirstName: "Sudhir",
		LastName:  "Phogat",
	}

	log.Println(user.FirstName, user.LastName, "Birthday=", user.DateOfBirth)
}
