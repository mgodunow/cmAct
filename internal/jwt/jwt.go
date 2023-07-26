package jwt

import (
	"cmAct/internal/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	Email string `json:"username"`
}

var (
	secretPhrase = utils.ReadSecret("secret_phrase")
)

func GenerateToken(email string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		email,
	})
	return t.SignedString([]byte(secretPhrase))
}
