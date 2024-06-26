package dtos

import "fmt"

type TransferRequest struct {
	FromAccountId uint64  `json:"from_account_id"`
	ToAccountId   uint64  `json:"to_account_number"`
	Amount        float64 `json:"amount"`
}

func (r *TransferRequest) ValidateAmount() error {
	if r.Amount < 0 {
		return fmt.Errorf("amount must be greater than zero\n")
	}
	return nil
}
