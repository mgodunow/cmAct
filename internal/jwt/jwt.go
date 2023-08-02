package jwt

import (
	"cmAct/internal/models"
	"cmAct/internal/utils"
	"errors"
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

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretPhrase), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *token.Claims")
	}
	acc, err := models.GetAccountByEmail(claims.Email)
	if err != nil {
		return "", err
	}

	return acc.Username, nil
}
