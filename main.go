package main

import (
	"fmt"
	"net/http"
)

//var can be changed so we define it as constant
//var portNumber = ":8080"

const portNumber = ":8080"

//we can move the responsewriter outside the main funciton

//in case there are mulitple pages then
// moved func home and about and render to handler

func main() {

	//***the handlefunc should be called before listenandserve else it will show the 404 error
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//Handle func is used to return response on request
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "Hello, World")
	//	//in case there is an error
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//Sprintf returns string where as n is an integer and cannot we writtern in println.
	//fmt.Printf("Number of Byte writtern are: %d", n)
	//})

	//in case I want to change the port number every time
	//_ = http.ListenAndServe(":8080", nil)

	fmt.Printf("Application starting at port number %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
