package main

import (
	"WebDevelopment/helpers"
	"fmt"
)

//creating packages
// in the terminal create a module by running the command "go mod init github.com/<yourid>/<modulename>"
//yourid just like mine "sudhirphogat" and modulename anythink like you want

//go.mod file is created with "module go mod init github.com/<yourid>/<modulename>""
func main() {

	var myVar helpers.SomeType
	myVar.TypeName = "This is taken from another folder"

	fmt.Println(myVar.TypeName)

}
