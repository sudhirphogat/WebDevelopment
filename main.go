package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Handle func is used to return response on request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World")
		//in case there is an error
		if err != nil {
			fmt.Println(err)
		}

		//Sprintf returns string where as n is an integer and cannot we writtern in println.
		fmt.Printf("Number of Byte writtern are: %d", n)
	})
	_ = http.ListenAndServe(":8080", nil)
}
