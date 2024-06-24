package scripts

import (
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"
	"github.com/Fernando-Balieiro/gobank/internal/infra/db"
	"log"
)

func seedAccount(fname, lname, passwd string, storage db.Storage) *domain.Account {
	acc, err := domain.NewAccount(fname, lname, passwd)
	if err != nil {
		log.Fatalln(err)
	}

	if err := storage.CreateAccount(acc); err != nil {
		log.Fatalln(err)
	}

	log.Printf("new account number => %d\n", acc.Number)

	return acc
}
func SeedAccounts(s db.Storage) {
	seedAccount("John", "Doe", "strongPasswd123", s)
	seedAccount("Jane", "Doe", "strongPasswd456", s)
	seedAccount("Foo", "Bar", "strongPasswd789", s)
}
