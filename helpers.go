package main

import (
	"math/rand"
	"time"
)

type gender []string

func getRandomGender() string {
	rand.Seed(time.Now().Unix())

	genders := gender{"Female", "Male", "Other"}
	randomGender := genders[rand.Intn(len(genders))]

	return randomGender
}
