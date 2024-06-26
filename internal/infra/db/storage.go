package db

import (
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"
)

type Storage interface {
	CreateAccount(domain *domain.Account) error
	DeleteAccount(id uint64) error
	GetAccountByID(id uint64) (*domain.Account, error)
	GetAccounts(searchQuery, sort string, limit, page int) ([]*domain.Account, error)
	GetAccountByNumber(number int) (*domain.Account, error)
	TransferMoney(idFrom, idTo uint64, balance float64) error
}
