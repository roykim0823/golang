package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var s2 = "six"

	s := "eight" // new local variable, not the global one

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething(("xxx"))

	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomthing func is", s)
	return s3, "World"
}
