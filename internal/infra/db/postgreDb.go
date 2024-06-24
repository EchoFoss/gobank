package db

import (
	"database/sql"
	"fmt"
	"github.com/Fernando-Balieiro/gobank/internal/domain"
	_ "github.com/lib/pq"
)

type PostgreDb struct {
	db *sql.DB
}

func (pg *PostgreDb) CreateAccount(account *domain.Account) error {
	query := `insert into accounts (first_name, last_name, balance, number, created_at)
    values ($1, $2, $3, $4, $5);`

	_, err := pg.db.Query(
		query,
		account.FirstName,
		account.LastName,
		account.Balance,
		account.Number,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil

}

func (pg *PostgreDb) DeleteAccount(id uint64) error {
	query := `delete from accounts where id = $1`
	_, err := pg.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("impossible to delete account with id: %d", id)
	}

	return nil
}

func (pg *PostgreDb) UpdateAccount(account *domain.Account) error {
	//TODO implement me
	panic("implement me")
}

func (pg *PostgreDb) GetAccountByID(id uint64) (*domain.Account, error) {
	query := `select * from accounts where id = $1`

	rows, err := pg.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account with id %d not found", id)
}

func (pg *PostgreDb) GetAccounts() ([]*domain.Account, error) {
	rows, err := pg.db.Query(`select * from accounts;`)

	if err != nil {
		return nil, err
	}

	var accounts []*domain.Account
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
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

func (pg *PostgreDb) Init() error {
	return pg.CreateAccountTable()
}
func (pg *PostgreDb) CreateAccountTable() error {
	query :=
		`create table if not exists accounts
			(
				id serial primary key,
				first_name varchar(50),
				last_name varchar(50),
				number serial,
				balance int,
				created_at timestamptz
			);`

	_, err := pg.db.Exec(query)
	return err
}

func scanIntoAccount(rows *sql.Rows) (*domain.Account, error) {
	account := new(domain.Account)
	err := rows.Scan(
		&account.Id,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}
