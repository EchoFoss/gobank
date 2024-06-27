package api

import (
	"encoding/json"
	"fmt"
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/login"
	"github.com/Fernando-Balieiro/gobank/internal/infra/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type WebServer struct {
	listenAddr string
	Storage    db.Storage
}

func NewWebServer(listenAddr string, storage db.Storage) *WebServer {
	return &WebServer{
		listenAddr: listenAddr,
		Storage:    storage,
	}
}

func (s *WebServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/login",
		makeHttpHandleFunc(s.handleLogin)).Methods(http.MethodPost)

	router.HandleFunc("/hello",
		SayHello).Methods(http.MethodGet)

	router.HandleFunc("/accounts",
		makeHttpHandleFunc(s.handleAccounts))

	router.HandleFunc("/accounts/{id}",
		withJWTAuth(makeHttpHandleFunc(s.handleAccountById), s.Storage)).Methods(http.MethodGet)

	log.Printf("API running on port %s", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Panic("erro ao iniciar api server")
	}
}

func (s *WebServer) handleAccounts(wr http.ResponseWriter, req *http.Request) error {

	if req.Method == http.MethodGet {
		return s.handleGetAccounts(wr, req)
	}
	if req.Method == http.MethodPost {
		return s.HandleCreateAccount(wr, req)
	}

	return fmt.Errorf("method not allowed: %s", req.Method)
}

func (s *WebServer) handleAccountById(wr http.ResponseWriter, req *http.Request) error {
	if req.Method == http.MethodGet {
		return s.handleGetAccountById(wr, req)
	}
	if req.Method == http.MethodDelete {
		return s.handleDeleteAccount(wr, req)
	}

	return fmt.Errorf("method not allowed: %s", req.Method)
}

func (s *WebServer) handleLogin(rw http.ResponseWriter, req *http.Request) error {
	if req.Method != http.MethodPost {
		return fmt.Errorf("method not allowed: %s", req.Method)
	}

	var logreq domain.LoginRequest
	if err := json.NewDecoder(req.Body).Decode(&logreq); err != nil {
		return err
	}

	acc, err := s.Storage.GetAccountByNumber(logreq.Number)
	if err != nil {
		return err
	}

	notAuthenticatedError := fmt.Errorf("not authenticated")

	if !acc.PasswordMatches(logreq.Password) {
		/*	return WriteJSON(rw, http.StatusForbidden, map[string]string{
			"login": "permission denied",
		}) */
		return notAuthenticatedError
	}

	if acc.Number != logreq.Number {
		/*return WriteJSON(rw, http.StatusForbidden, map[string]string{
			"login": "permission denied",
		})*/
		return notAuthenticatedError
	}

	token, err := createJWT(acc)
	if err != nil {
		return err
	}

	resp := domain.LoginResponse{
		Number: acc.Number,
		Token:  token,
	}

	return WriteJSON(rw, http.StatusOK, resp)
}

func WriteJSON(wr http.ResponseWriter, status int, v any) error {
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(status)

	return json.NewEncoder(wr).Encode(v)
}

type ErrorAPI struct {
	Error string `json:"error"`
}

type apiFunc func(wr http.ResponseWriter, req *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		if err := f(wr, req); err != nil {
			err := WriteJSON(wr, http.StatusBadRequest, ErrorAPI{err.Error()})
			if err != nil {
				return
			}
		}
	}
}
