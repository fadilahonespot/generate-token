package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	token, err := CreateToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INI TOKEN: ", token)
}

func CreateToken() (string, error) {
	var err error
	//Creating Access Token
	key := "testing"

	token := jwt.New(jwt.SigningMethodHS256)

	// Bikin payload
	claims := token.Claims.(jwt.MapClaims)
	claims["access_id"] = "db855566-2afd-11ec-8c7c-0242ac130028"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["id"] = "645ef910-2aeb-11ec-8c7c-0242ac130065"
	claims["permissions"] = nil
	claims["role"] = "Super user"

	// Di tandatangani
	t, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}


	return t, nil
}