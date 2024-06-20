package api

import (
	"encoding/json"
	"fmt"
	"github.com/Fernando-Balieiro/gobank/internal/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// WriteJSON TODO: O header não está retornando application/json na resposta da requisição
func WriteJSON(wr http.ResponseWriter, status int, v any) error {
	wr.WriteHeader(status)
	fmt.Println(wr.Header())
	wr.Header().Set("Content-Type", "application/json")
	fmt.Println(wr.Header())

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

type WebServer struct {
	listenAddr string
}

func NewWebServer(listenAddr string) *WebServer {
	return &WebServer{listenAddr: listenAddr}
}

func (s *WebServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))

	log.Printf("API running on port %s", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Panic("erro ao iniciar api server")
	}
}

func (s *WebServer) handleAccount(wr http.ResponseWriter, req *http.Request) error {

	if req.Method == "GET" {
		return s.handleGetAccount(wr, req)
	}
	if req.Method == "POST" {
		return s.handleGetAccount(wr, req)
	}
	if req.Method == "DELETE" {
		return s.handleGetAccount(wr, req)
	}

	return fmt.Errorf("method not allowed: %s", req.Method)
}

func (s *WebServer) handleGetAccount(wr http.ResponseWriter, req *http.Request) error {
	account := domain.NewAccount("Fernando", "Balieiro")
	return WriteJSON(wr, http.StatusOK, account)
}

func (s *WebServer) handleCreateAccount(wr http.ResponseWriter, req *http.Request) error {
	return nil
}
func (s *WebServer) handleDeleteAccount(wr http.ResponseWriter, req *http.Request) error {
	return nil
}
func (s *WebServer) handleTransfer(wr http.ResponseWriter, req *http.Request) error {
	return nil
}
