package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//****JSON
//*** if we do not know the fields in the struct or Json that there is another method that is not covered here
//use `` in ~ button and not '' in " " button.
//json:"name used by json" to identify the value of variable
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `jsaon:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name":"Sudhir",
		"last_name":"Phogat",
		"hair_color":"Black",
		"has_dog":true
	},
	{
		"first_name":"Riaan",
		"last_name":"Phogat",
		"hair_color":"Brown",
		"has_dog":false
	}
]`
	//Json uses marshalled and unmarshalled

	//creating struct from json
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error Unmarshalled Json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	//creating Json from struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Renu"
	m1.LastName = "Kadyan"
	m1.HairColor = "Black"
	m1.HasDog = true

	var m2 Person
	m2.FirstName = "Renu"
	m2.LastName = ""
	m2.HairColor = "Black"
	m2.HasDog = true

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)
	//instead of json.marshall I can give marshalintend so that we can product neat output
	//but in production we use marshall
	//newJson, err := json.Marshal(mySlice)
	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("Error in creating Json", err)
	}
	fmt.Println(string(newJson))
}
