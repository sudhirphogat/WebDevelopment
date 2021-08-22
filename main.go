package main

import "fmt"

//Interfaces
///type interface with methods
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	name  string
	breed string
}
type Gorilla struct {
	name          string
	numberOfTeeth int
	color         string
}

func main() {

	dog := Dog{
		name:  "Tommy",
		breed: "German",
	}
	//dog does not suffice the requiremnets so we create another 2 functions fro say and number of legs
	// add & because we are passing the pointer reference to the below function
	fmt.Println(dog.name)
	printInfo(&dog)

	//the value of name can be changed in the below function where I pass the pointer value
	fmt.Println(dog.name)

	gorilla := Gorilla{
		//	name:          "GoriallaName",
		///	color:         "blue",
		//numberOfTeeth: 45,
	}
	//if we add & while sending value to function, the function can take pointers or value based on what we give
	printInfo(&gorilla)

}

// The values defined above in main function to gorilla and dog are useless
//the entire output is printed based on below 2 functions
func printInfo(a Animal) {

	fmt.Println("the animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")

}

//we can either pass the type here or the pointers
//pointer is best practise by adding * to the type
func (d *Dog) Says() string {
	d.name = "diky"
	return "woof"
}

// for learning and testing purpose I passed the type here and did not add * but it works
func (d Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "haha"
}

func (d *Gorilla) NumberOfLegs() int {
	return 5
}
