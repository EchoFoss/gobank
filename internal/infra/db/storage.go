package db

import domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"

type Storage interface {
	CreateAccount(domain *domain.Account) error
	DeleteAccount(id uint64) error
	GetAccountByID(id uint64) (*domain.Account, error)
	GetAccounts() ([]*domain.Account, error)
	GetAccountbyNumnber(number int) (*domain.Account, error)
}
