package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"eamil"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Address
	Company
}

type response []user

func getUsers() []user {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	var result response
	e := json.NewDecoder(resp.Body).Decode(&result) // response body is []byte

	if e != nil {
		log.Fatal("could not read body")
	}

	return result
}
