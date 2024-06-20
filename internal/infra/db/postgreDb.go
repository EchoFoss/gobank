package db

import (
	"database/sql"
	"github.com/Fernando-Balieiro/gobank/internal/domain"
	_ "github.com/lib/pq"
)

type PostgreDb struct {
	db *sql.DB
}

func (p *PostgreDb) CreateAccount(account *domain.Account) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgreDb) DeleteAccount(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgreDb) UpdateAccount(account *domain.Account) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgreDb) GetAccountByID(id uint64) (*domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostgreDb() (*PostgreDb, error) {
	connStr := "user=admin dbname=gobank password=passwd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreDb{
		db: db,
	}, nil
}
