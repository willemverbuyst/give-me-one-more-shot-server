package main

import (
	"errors"

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

func createPatient(userName string) patient {
	active := true
	birthDate := "12/12/2012"
	bsn := getRandomBSN()
	email := "a@a.com"
	familyName := userName
	gender := getRandomGender()
	givenName := "BBB"
	id := uuid.New().String()

	p := patient{active, birthDate, bsn, email, familyName, gender, givenName, id}
	return p
}

func createDummyPatients() []patient {
	users := getUsers()
	dummyPatients := []patient{}
	for i := 1; i < 10; i++ {
		userName := users[i].Name
		dummyPatients = append(dummyPatients, createPatient(userName))
	}
	return dummyPatients
}

func getPatientById(id string) (*patient, error) {
	for i, p := range patients {
		if p.ID == id {
			return &patients[i], nil
		}
	}
	return nil, errors.New("patient not found")
}
