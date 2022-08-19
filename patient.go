package main

import (
	"errors"
	"strings"

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

func createPatient(f string, g string, e string) patient {
	active := true
	birthDate := getRandomDate()
	bsn := getRandomBSN()
	email := e
	familyName := f
	gender := getRandomGender()
	givenName := g
	id := uuid.New().String()

	p := patient{active, birthDate, bsn, email, familyName, gender, givenName, id}
	return p
}

func createDummyPatients() []patient {
	users := getUsers()

	dummyPatients := []patient{}
	for i := 1; i < 10; i++ {
		name := removePrefixFromName(users[i].Name)
		familyName := strings.Split(name, " ")[0]
		givenName := strings.Join(strings.Split(name, " ")[1:], " ")
		email := createEmailWithName(name)
		dummyPatients = append(dummyPatients, createPatient(familyName, givenName, email))
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
