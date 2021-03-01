package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)


func CreateToken(username string) (string,error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() 

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "Error", err
	}

	return t, err
}