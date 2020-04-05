package services

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("todokey")

type myCustomclaims struct {
	userID int `json.userid`
	jwt.StandardClaims
}

//GenerateToken jwt
func GenerateToken(userID int) (string, error) {
	//create JWT claims

	claims := jwt.MapClaims{}
	claims["userId"] = userID
	//declare token with algorithm used for signing and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//create jwtString
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

//VerifyToken handler
func VerifyToken(tokenString string) (int, error) {
	// Initialize a new instance of `Claims`

	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil && !tkn.Valid {
		log.Fatal(err.Error())
		return 0, errors.New("Unauthorized")
	}

	// userID := claims["userId"]
	id := claims["userId"].(float64)
	return int(id), nil
}
