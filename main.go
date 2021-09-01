package main

import (
	"errors"
	"fmt"
	"net/http"
)

//var can be changed so we define it as constant
//var portNumber = ":8080"

const portNumber = ":8080"

//we can move the responsewriter outside the main funciton

//in case there are mulitple pages then

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home Page")

}

//About is called when /about is called
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 5)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("The total add values is %d", sum))
}

// simple example to return the sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100, 0)
	if err != nil {
		fmt.Fprintf(w, "Cannot print by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("the divide answer is %f", f))

}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by Zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {

	//***the handlefunc should be called before listenandserve else it will show the 404 error
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

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
