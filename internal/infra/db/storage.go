package db

import "github.com/Fernando-Balieiro/gobank/internal/domain"

type Storage interface {
	CreateAccount(account *domain.Account) error
	DeleteAccount(id uint64) error
	UpdateAccount(*domain.Account) error
	GetAccountByID(id uint64) (*domain.Account, error)
}
