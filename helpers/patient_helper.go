package helpers

import (
	"errors"
	"math/rand"
	"server/models"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type gender []string

func getRandomGender() string {
	rand.Seed(time.Now().UnixNano())

	genders := gender{"Female", "Male", "Other"}
	randomGender := genders[rand.Intn(len(genders))]

	return randomGender
}

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10)

	return randomNumber
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

func getRandomBSN() string {
	var BSN = models.BSN{
		Leading: 9999,
		I4:      9,
		I5:      getRandomNumber(),
		I6:      getRandomNumber(),
		I7:      getRandomNumber(),
		I8:      0,
	}

	sum := getSum(BSN)
	BSN.I8 = sum - ((sum / 11) * 11)

	if BSN.I8 > 9 && BSN.I7 > 0 {
		BSN.I7 = BSN.I7 - 1
		BSN.I8 = 8
	}

	if BSN.I8 > 9 && !(BSN.I7 > 0) {
		BSN.I7 = BSN.I7 + 1
		BSN.I8 = 1
	}

	return strconv.Itoa(BSN.Leading) + strconv.Itoa(BSN.I4) + strconv.Itoa(BSN.I5) + strconv.Itoa(BSN.I6) + strconv.Itoa(BSN.I7) + strconv.Itoa(BSN.I8)
}

func getSum(BSN models.BSN) int {
	sum := 0
	for i, num := range splitInt(BSN.Leading) {
		sum = sum + num*(9-i)
	}

	return sum + 5*BSN.I4 + 4*BSN.I5 + 3*BSN.I6 + 2*BSN.I7
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
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

	return strings.ToLower(email)
}

func createPatient(f string, g string, e string) models.Patient {
	active := true
	birthDate := getRandomDate()
	bsn := getRandomBSN()
	email := e
	familyName := f
	gender := getRandomGender()
	givenName := g
	id := uuid.New().String()

	p := models.Patient{
		Active:     active,
		Birthdate:  birthDate,
		BSN:        bsn,
		Email:      email,
		FamilyName: familyName,
		Gender:     gender,
		GivenName:  givenName,
		ID:         id,
	}
	return p
}

func CreateDummyPatients(users []models.User) []models.Patient {
	dummyPatients := []models.Patient{}
	for i := 1; i < 10; i++ {
		name := removePrefixFromName(users[i].Name)
		familyName := strings.Split(name, " ")[0]
		givenName := strings.Join(strings.Split(name, " ")[1:], " ")
		email := createEmailWithName(name)
		dummyPatients = append(dummyPatients, createPatient(familyName, givenName, email))
	}
	return dummyPatients
}

func GetPatientById(id string, patients []models.Patient) (*models.Patient, error) {
	for i, p := range patients {
		if p.ID == id {
			return &patients[i], nil
		}
	}
	return nil, errors.New("patient not found")
}
