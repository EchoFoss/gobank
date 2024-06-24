package domain_test

import (
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc, err := domain.NewAccount("Fernando", "Balieiro", "passwdTeste")
	assrt := assert.New(t)
	assrt.NotNil(acc)
	assrt.Nil(err)
}
