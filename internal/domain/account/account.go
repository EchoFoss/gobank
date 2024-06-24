package domain

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// Account Usar apenas a função de fábrica para iniciar essa struct
type Account struct {
	Id                uint64    `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Number            int       `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           float64   `json:"balance"`
	CreatedAt         time.Time `json:"created_at"`
}

func (a *Account) PasswordMatches(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encriptedPassWd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		Number:            int(rand.Int63n(10000)),
		EncryptedPassword: string(encriptedPassWd),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
