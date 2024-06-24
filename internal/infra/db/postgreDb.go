package db

import (
	"database/sql"
	"fmt"
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"
	_ "github.com/lib/pq"
	"log"
)

type PostgreDb struct {
	db *sql.DB
}

func (pg *PostgreDb) CreateAccount(account *domain.Account) error {
	query := `insert into accounts (first_name, last_name, balance, encrypted_password, number, created_at)
    values ($1, $2, $3, $4, $5, $6);`

	_, err := pg.db.Query(
		query,
		account.FirstName,
		account.LastName,
		account.Balance,
		account.EncryptedPassword,
		account.Number,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil

}

func (pg *PostgreDb) GetAccountbyNumnber(number int) (*domain.Account, error) {
	query := `select * from accounts where number = $1;`
	rows, err := pg.db.Query(query, number)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with id [%d] not found", number)
}

func (pg *PostgreDb) DeleteAccount(id uint64) error {
	query := `delete from accounts where id = $1`
	_, err := pg.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("impossible to delete account with id: %d", id)
	}

	return nil
}

// Possivelmente não será implementado, já que uma conta de banco não tem nada além do saldo atualizado
//func (pg *PostgreDb) UpdateAccount(account *domain.Account) error {
//	panic("implement me")
//}

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

/*
Init TODO: criar o notification pattern para retornar todos os erros do banco de dados, caso algum aconteça
*/
func (pg *PostgreDb) Init() error {
	err := pg.createAccountTable()
	if err != nil {
		log.Printf("error initializing accounts table if it doesnt exist: %+v", err)
		return err
	}
	return nil
}
func (pg *PostgreDb) createAccountTable() error {

	query :=
		`create table if not exists accounts
			(
				id serial primary key,
				first_name varchar(50),
				last_name varchar(50),
				number int,
    			encrypted_password text,
				balance int,
				created_at timestamptz
			);`

	_, err := pg.db.Exec(query)
	return err
}

func scanIntoAccount(rows *sql.Rows) (*domain.Account, error) {
	account := domain.Account{}
	err := rows.Scan(
		&account.Id,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt,
	)

	return &account, err
}
