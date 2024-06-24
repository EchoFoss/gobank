package domain

import (
	"math/rand"
	"time"
)

type Account struct {
	Id        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int       `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int(rand.Int63n(10000)),
		CreatedAt: time.Now().UTC(),
	}
}
