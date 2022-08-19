package main

import (
	"errors"
	"math/rand"
	"strings"
	"time"

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

func removePrefixFromName(name string) string {
	elements := strings.Split(name, " ")

	for i, e := range elements {
		if strings.HasSuffix(e, ".") {
			elements = append(elements[:i], elements[i+1:]...)
		}
	}

	return strings.Join(elements, " ")
}

func createEmailWithName(name string) string {
	elements := strings.Split(name, " ")
	suffix := getRandomEmailSuffix()
	email := elements[0] + "@" + elements[1] + "." + suffix

	return email
}

func getRandomEmailSuffix() string {
	rand.Seed(time.Now().UnixNano())

	s := gender{"io", "com", "org"}
	randomSuffix := s[rand.Intn(len(s))]

	return randomSuffix
}

func getRandomDate() string {
	min := time.Date(1950, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}
