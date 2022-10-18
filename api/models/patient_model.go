package models

type Patient struct {
	Active     bool   `json:"active" example:"true"`
	Birthdate  string `json:"birthdate" exmaple:"2012-12-12"`
	BSN        string `json:"BSN" example:"999999999"`
	Email      string `json:"email" example:"a@b.io"`
	FamilyName string `json:"familyName" example:"Doe"`
	Gender     string `json:"gender" example:"Female"`
	GivenName  string `json:"givenName" example:"Jane"`
	ID         string `json:"id" example:"46cb51b9-68ae-4560-943a-8ea0ae3ddc14"`
}
