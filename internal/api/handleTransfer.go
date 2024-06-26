package api

import (
	"encoding/json"
	"github.com/Fernando-Balieiro/gobank/internal/domain/dtos"
	"log"
	"net/http"
)

func (s *WebServer) handleTransfer(wr http.ResponseWriter, req *http.Request) error {
	transferReq := dtos.TransferRequest{}

	if err := json.NewDecoder(req.Body).Decode(&transferReq); err != nil {
		return err
	}

	defer func() {
		err := req.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	if err := transferReq.ValidateAmount(); err != nil {
		return err
	}

	err := s.Storage.TransferMoney(transferReq.FromAccountId, transferReq.ToAccountId, transferReq.Amount)

	if err != nil {
		return WriteJSON(wr, http.StatusBadRequest, map[string]string{
			"error transfering money": err.Error(),
		})
	}

	return WriteJSON(wr, http.StatusOK, "transfer succeeded")
}
