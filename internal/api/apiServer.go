package api

import (
	"encoding/json"
	"fmt"
	"github.com/Fernando-Balieiro/gobank/internal/domain"
	"github.com/Fernando-Balieiro/gobank/internal/domain/dtos"
	"github.com/Fernando-Balieiro/gobank/internal/infra/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type WebServer struct {
	listenAddr string
	storage    db.Storage
}

func NewWebServer(listenAddr string, storage db.Storage) *WebServer {
	return &WebServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func (s *WebServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", makeHttpHandleFunc(helloWord)).Methods(http.MethodGet)

	router.HandleFunc("/accounts", makeHttpHandleFunc(s.handleAccount))
	router.HandleFunc("accounts/{accountId}", makeHttpHandleFunc(s.handleGetAccountById)).Methods(http.MethodGet)

	log.Printf("API running on port %s", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Panic("erro ao iniciar api server")
	}
}

func (s *WebServer) handleAccount(wr http.ResponseWriter, req *http.Request) error {

	if req.Method == http.MethodGet {
		return s.handleGetAccounts(wr, req)
	}
	if req.Method == http.MethodPost {
		return s.handleCreateAccount(wr, req)
	}
	if req.Method == http.MethodDelete {
		return s.handleDeleteAccount(wr, req)
	}

	return fmt.Errorf("method not allowed: %s", req.Method)
}

func (s *WebServer) handleGetAccountById(wr http.ResponseWriter, req *http.Request) error {
	id := mux.Vars(req)["accountId"]

	fmt.Println(id)
	return WriteJSON(wr, http.StatusOK, &domain.Account{})
}

// GET /accounts
func (s *WebServer) handleGetAccounts(wr http.ResponseWriter, req *http.Request) error {
	accounts, err := s.storage.GetAccounts()

	if err != nil {
		return err
	}
	return WriteJSON(wr, http.StatusOK, &accounts)
}

func (s *WebServer) handleCreateAccount(wr http.ResponseWriter, req *http.Request) error {
	createAccountReq := new(dtos.CreateAccountDto)
	if err := json.NewDecoder(req.Body).Decode(createAccountReq); err != nil {
		return err
	}

	account := domain.NewAccount(createAccountReq.FirstName, createAccountReq.LastName)

	if err := s.storage.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(wr, http.StatusCreated, account)

}
func (s *WebServer) handleDeleteAccount(wr http.ResponseWriter, req *http.Request) error {
	return nil
}
func (s *WebServer) handleTransfer(wr http.ResponseWriter, req *http.Request) error {
	return nil
}

func WriteJSON(wr http.ResponseWriter, status int, v any) error {
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(status)

	return json.NewEncoder(wr).Encode(v)
}

type ApiError struct {
	Error string
}

type apiFunc func(wr http.ResponseWriter, req *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		if err := f(wr, req); err != nil {
			WriteJSON(wr, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

func helloWord(wr http.ResponseWriter, req *http.Request) error {
	_, _ = wr.Write([]byte("hello word"))
	return nil
}
