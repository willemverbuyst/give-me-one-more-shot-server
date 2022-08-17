package main

import (
	"math/rand"
	"strconv"
	"time"
)

type gender []string

type bsn struct {
	leading int
	i4      int
	i5      int
	i6      int
	i7      int
	i8      int
}

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

func getRandomBSN() string {
	var BSN = bsn{
		leading: 9999,
		i4:      9,
		i5:      getRandomNumber(),
		i6:      getRandomNumber(),
		i7:      getRandomNumber(),
		i8:      0,
	}

	sum := getSum(BSN)
	BSN.i8 = sum - ((sum / 11) * 11)

	if BSN.i8 > 9 && BSN.i7 > 0 {
		BSN.i7 = BSN.i7 - 1
		BSN.i8 = 8
	}

	if BSN.i8 > 9 && !(BSN.i7 > 0) {
		BSN.i7 = BSN.i7 + 1
		BSN.i8 = 1
	}

	return strconv.Itoa(BSN.leading) + strconv.Itoa(BSN.i4) + strconv.Itoa(BSN.i5) + strconv.Itoa(BSN.i6) + strconv.Itoa(BSN.i7) + strconv.Itoa(BSN.i8)
}

func getSum(BSN bsn) int {
	sum := 0
	for i, num := range splitInt(BSN.leading) {
		sum = sum + num*(9-i)
	}

	return sum + 5*BSN.i4 + 4*BSN.i5 + 3*BSN.i6 + 2*BSN.i7
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}
