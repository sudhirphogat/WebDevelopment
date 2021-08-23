package main

import (
	"WebDevelopment/helpers"
	"fmt"
	"log"
)

//the random number will be below numPool
const numPool = 10

//creating packages
// in the terminal create a module by running the command "go mod init github.com/<yourid>/<modulename>"
//yourid just like mine "sudhirphogat" and modulename anythink like you want

//go.mod file is created with "module go mod init github.com/<yourid>/<modulename>""
func main() {

	var myVar helpers.SomeType
	myVar.TypeName = "This is taken from another folder"

	fmt.Println(myVar.TypeName)

	//***Channels
	//*** channels are useful when we have more than many packages inside project

	//we can pass value to fucntions and get return value from them through channels
	//initialise the channel ** similar to creating MAP
	intChan := make(chan int)

	//close the channel as a good practise anyhow it is not required.
	//it is similar to closing a sql connection or file after it is used
	defer close(intChan)

	//concurrent operation by sending it as routine by using go
	go CalculateValue(intChan)

	//get the response from the channel after sending from above routine
	//<- is what is assigned to intchan
	num := <-intChan

	log.Println(num)
	//When running the program always the result is same 1.
	//To fix this we need to make the random number psudo in helper.

}

//creating function for the channal to create random number
//this will accept a channel value
//Random value will be created in Helper folder under helpers
func CalculateValue(intChan chan int) {

	//Function is called from helper and value is passed which is declared in package scope above
	randomNumber := helpers.RandomNumber(numPool)
	//pass the random number through channel. where intChan is the channel we created
	intChan <- randomNumber

}
