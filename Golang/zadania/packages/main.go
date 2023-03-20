package main

import (
	"github.com/google/uuid"
	"github.com/grupawp/appdispatcher"
	"log"
)

type Student struct {
	FirstName     string
	LastName      string
	applicationID uuid.UUID
}

func (s Student) ApplicationID() string {
	return s.applicationID.String()
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func main() {
	student := Student{
		FirstName:     "Paweł",
		LastName:      "Saniewski",
		applicationID: uuid.New(),
	}

	statusCode, err := appdispatcher.Submit(student)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(statusCode)
}
