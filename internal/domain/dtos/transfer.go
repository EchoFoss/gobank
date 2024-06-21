package dtos

type TransferRequest struct {
	ToAccountId uint64  `json:"to_account_number"`
	Amount      float32 `json:"amount"`
}
