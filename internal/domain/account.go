package domain

import "math/rand"

type Account struct {
	Id        uint64  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Number    uint64  `json:"number"`
	Balance   float64 `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:        uint64(rand.Intn(1000)),
		FirstName: firstName,
		LastName:  lastName,
		Number:    uint64(rand.Intn(1000000)),
	}
}
