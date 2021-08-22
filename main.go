package main

import (
	"log"
)

func main() {
	//loops and range
	//loops
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			log.Println(i)
		}
	}

	//**Range with slice

	animals := []string{"dog", "fish", "Riaan", "horse"}
	// i will print the index and animal will print one by one range in animal

	//for i, animal := range animals {
	// _ is used to keep the variable as blank and ignore so that Go do not throw an error
	for _, animal := range animals {
		//log.Println(i, animal)
		log.Println(animal)
	}

	//**loop range with Map

	animalsMap := make(map[string]string)
	animalsMap["dog"] = "charli"
	animalsMap["horse"] = "pinto"

	for _, animal := range animalsMap {
		log.Println(animal)
	}

	//if we want to print the animal type and their name
	for animalType, animal := range animalsMap {
		log.Println(animalType, animal)
	}

	//**loop range through string

	var firstLine = "This is a string"
	for i, l := range firstLine {
		log.Println(i, l)
		//this will not print the letter but the bytes

		log.Printf("%v\t%#U\t", i, l)
	}

	//**range over customer types

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var user []User
	user = append(user, User{"Sudhir", "Phogat", "sudhir@gmail.com", 30})
	user = append(user, User{"Ram", "Kumar", "ram@gmail.com", 90})
	user = append(user, User{"Siva", "Mani", "siva@gmail.com", 69})

	for _, u := range user {
		log.Println(u.FirstName, u.LastName, u.Email, u.Age)
	}

}
