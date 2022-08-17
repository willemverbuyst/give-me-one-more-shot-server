package main

type patient struct {
	Birthdate  string `json:"birthdate"`
	BSN        string `json:"BSN"`
	Email      string `json:"email"`
	FamilyName string `json:"familyName"`
	Gender     string `json:"gender"`
	GivenName  string `json:"givenName"`
}

func createPatient() string {
	gender := getRandomGender()
	bsn := getRandomBSN()

	return gender + "---" + bsn
}
