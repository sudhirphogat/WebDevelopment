package helpers

import (
	"math/rand"
	"time"
)

//I created a folder and in that folder i created this file
//package name above is folder name

//I created below struct here so I will use the same in main.go

type SomeType struct {
	TypeName   string
	TypeNumber int
}

//function for channel example to create random number
//function will take n integer and return integer
func RandomNumber(n int) int {

	//this rand.seed is added to make the random number unique, else always same number was generated
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
