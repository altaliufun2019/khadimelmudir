package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"mudiralmaham/utils/authentication"
	"time"
)

func jwtEncoder(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(authentication.SecretKey)
	return tokenString, err
}
