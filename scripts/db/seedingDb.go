package scripts

import (
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"
	"github.com/Fernando-Balieiro/gobank/internal/infra/db"
	"github.com/go-faker/faker/v4"
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
	for i := 0; i < 50; i++ {
		fname := faker.FirstName()
		lname := faker.LastName()
		password := faker.Password()
		usr, err := domain.NewAccount(fname, lname, password)
		if err != nil {
			log.Printf("erro ao criar usuário: %v\n", err)
		}
		go seedAccount(usr.FirstName, usr.LastName, usr.EncryptedPassword, s)
	}
}
