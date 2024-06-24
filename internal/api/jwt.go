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
			log.Printf("Erro ao validar jwt: %+v\n", err)
			return
		}

		// não sei se isso é redundante
		if !token.Valid {
			permissionDenied(wr)
			log.Printf("jwt não é valido: %+v\n", err)
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
			log.Printf("não foi possível pegar o ID do usuário em jwt.go: %+v\n", err)

			return
		}

		if err != nil {
			permissionDenied(wr)
			log.Printf("not abled to convert accountNumber to int")
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
		accountNumberParsed, err := strconv.Atoi(accountNumber)
		if err != nil {
			permissionDenied(wr)
			//log.Printf("%+v", err)
			//log.Println(claims["accountNumber"])
			return
		}
		if user.Number != accountNumberParsed {
			permissionDenied(wr)
			//log.Printf("JWT Account Number: %d", claims["accountNumber"])
			//log.Printf("Account number: %d", user.Number)
			return

		}
		// Mapa mental do que posso fazer
		//account, err := store.GetUserByJWTToken(token)
		//// /account/{id}
		//
		//claims := token.Claims.(jwt.MapClaims)
		//if claims["Id"] == account.Id {
		//
		//}
		//fmt.Println(claims)
		//fmt.Println(userID)
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
	expirationDate := jwt.NewNumericDate(time.Now().Add(time.Hour * 15))

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
