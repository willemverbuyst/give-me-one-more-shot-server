package main

import (
	"github.com/google/uuid"
)

type patient struct {
	Active     bool   `json:"active"`
	Birthdate  string `json:"birthdate"`
	BSN        string `json:"BSN"`
	Email      string `json:"email"`
	FamilyName string `json:"familyName"`
	Gender     string `json:"gender"`
	GivenName  string `json:"givenName"`
	ID         string `json:"id"`
}

func createPatient() patient {
	active := true
	birthDate := "12/12/2012"
	bsn := getRandomBSN()
	email := "a@a.com"
	familyName := "AAA"
	gender := getRandomGender()
	givenName := "BBB"
	id := uuid.New().String()

	p := patient{active, birthDate, bsn, email, familyName, gender, givenName, id}
	return p
}

func createDummyPatients() []patient {
	dummyPatients := []patient{}
	for i := 1; i < 11; i++ {
		dummyPatients = append(dummyPatients, createPatient())
	}
	return dummyPatients
}
