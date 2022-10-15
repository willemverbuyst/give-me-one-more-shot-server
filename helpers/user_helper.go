package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/models"
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
