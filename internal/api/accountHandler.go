package api

import (
	"encoding/json"
	"fmt"
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"
	"github.com/Fernando-Balieiro/gobank/internal/domain/dtos"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// HandleCreateAccount POST /accounts
func (s *WebServer) HandleCreateAccount(wr http.ResponseWriter, req *http.Request) error {
	createAccountReq := dtos.CreateAccountDto{}
	if err := json.NewDecoder(req.Body).Decode(&createAccountReq); err != nil {
		return err
	}

	account, err := domain.NewAccount(createAccountReq.FirstName, createAccountReq.LastName, createAccountReq.Password)
	if err != nil {
		return err
	}

	if err := s.Storage.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(wr, http.StatusCreated, account)

}

// HandleGetAccounts GET /accounts
func (s *WebServer) handleGetAccounts(wr http.ResponseWriter, req *http.Request) error {
	searchQuery := req.URL.Query()

	search := searchQuery.Get("search")
	mapErrs := make(map[string]string)

	sort := searchQuery.Get("sort")

	limit := searchQuery.Get("limit")

	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		mapErrs["limit is not a number"] = "Limit must be an integer"
	}

	page := searchQuery.Get("page")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		mapErrs["page is not a number"] = "Page must be an integer"
	}

	if len(mapErrs) > 0 && mapErrs != nil {
		return WriteJSON(wr, http.StatusBadRequest, mapErrs)
	}

	accounts, err := s.Storage.GetAccounts(search, sort, limitNum, pageNum)

	if err != nil {
		return err
	}
	return WriteJSON(wr, http.StatusOK, &accounts)
}

// HandleGetAccountById GET /accounts/{id}
func (s *WebServer) handleGetAccountById(wr http.ResponseWriter, req *http.Request) error {
	id, err := getId(req)
	if err != nil {
		return err
	}
	account, err := s.Storage.GetAccountByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(wr, http.StatusOK, account)
}

// handleDeleteAccount DELETE /accounts
func (s *WebServer) handleDeleteAccount(wr http.ResponseWriter, req *http.Request) error {

	id, err := getId(req)
	if err != nil {
		return err
	}
	if err := s.Storage.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJSON(wr, http.StatusOK, map[string]uint64{"deleted": id})
}

func getId(r *http.Request) (uint64, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return id, fmt.Errorf("invalid id given: %s", idStr)
	}
	return id, nil
}
