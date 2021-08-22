package main

import "log"

//decision if & switch case
func main() {
	var myVar string
	myVar = "dog"
	isTrue := false

	if myVar == "cat" && isTrue {
		log.Println("The variable is cat and isTrue is true")
	} else if myVar == "dog" && isTrue == false {
		log.Println("The var is dog and isTrue is false")
	} else if myVar == "cat" || isTrue {
		log.Println("This is an or statement of cat and true")
	} else {
		log.Println("nothing is correct")
	}

	myVar2 := "sudhir"

	switch myVar2 {
	case "sudhir":
		log.Println("THis is case 1")

	case "Sudhir":
		log.Println("this is case 2")

	default:
		log.Println("THi is default")

	}

}
