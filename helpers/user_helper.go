package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"give-me-one-more-shot/server/models"
	"log"
	"net/http"
	"strconv"
)

type response []models.User

func GetUsers() response {
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

func GetUserById(id string, users []models.User) (*models.User, error) {
	for i, u := range users {
		idStr, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Conversion of id did not work")
		}
		if u.ID == idStr {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}
