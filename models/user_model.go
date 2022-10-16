package models

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	}
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

type User struct {
	Active   bool   `json:"active"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"eamil"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Address
	Company
}

type BSN struct {
	Leading int `json:"leading"`
	I4      int `json:"i4"`
	I5      int `json:"i5"`
	I6      int `json:"i6"`
	I7      int `json:"i7"`
	I8      int `json:"i8"`
}
