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
	return WriteJSON(wr, http.StatusOK, transferReq)
}
