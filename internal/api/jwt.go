package api

import (
	"fmt"
	domain "github.com/Fernando-Balieiro/gobank/internal/domain/account"
	"github.com/Fernando-Balieiro/gobank/internal/infra/db"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

func permissionDenied(w http.ResponseWriter) {
	_ = WriteJSON(w, http.StatusForbidden, ErrorAPI{"permission denied"})
}

// TODO: Implementar SRP nessa função para desacoplar a funcionalidade dela
func withJWTAuth(handlerFunc http.HandlerFunc, store db.Storage) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		fmt.Println("calling JWT auth middleware")

		tokenString := req.Header.Get("x-jwt-token")

		token, err := validateJWT(tokenString)

		if err != nil {
			permissionDenied(wr)
			return
		}

		// não sei se isso é redundante
		if !token.Valid {
			permissionDenied(wr)
			return
		}

		userID, err := getId(req)
		if err != nil {
			_ = WriteJSON(wr, http.StatusForbidden, ErrorAPI{Error: "erro ao parsear o id"})
			return
		}

		user, err := store.GetAccountByID(userID)
		if err != nil {
			//permissionDenied(wr)
			permissionDenied(wr)

			return
		}

		if err != nil {
			permissionDenied(wr)
			return
		}

		/*
			horas perdidas aqui: 5
			atualizar conforme foram sendo perdidas mais horas

			O erro que foi descobrido que estava dando era um ao pegar a secret, qualquer erro primeiramente
			procurar na funcão de dar um get na secret
		*/
		claims := token.Claims.(jwt.MapClaims)

		accountNumber := claims["accountNumber"].(string)
		expiresAtTime, err := time.Parse(time.UnixDate, claims["expiresAt"].(string))

		if expiresAtTime.Before(time.Now()) {
			permissionDenied(wr)
			return
		}

		if err != nil {
			log.Println("error loading timezone")
		}

		accountNumberParsed, err := strconv.Atoi(accountNumber)
		if err != nil {
			permissionDenied(wr)
			return
		}
		if user.Number != accountNumberParsed {
			permissionDenied(wr)
			return

		}

		handlerFunc(wr, req)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	// TODO: setar um .env corretamente para o secret jwt
	secret := getJWTSecret()
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func getJWTSecret() string {
	return "gobank123"
	//return os.Getenv("JWT_SECRET")
}

func createJWT(account *domain.Account) (string, error) {
	expirationDate := jwt.NewNumericDate(time.Now().Add(time.Hour * 15).UTC())

	claims := jwt.MapClaims{
		"expiresAt": expirationDate,

		/*
			HACK(Fernando): accountNumber é parseado como float64 de forma default
			parsear para string depois
		*/
		"accountNumber": strconv.Itoa(account.Number),
	}

	secret := getJWTSecret()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
