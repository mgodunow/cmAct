package hash

import (
	"crypto/sha1"
	"fmt"

	"cmAct/internal/utils"
)

var salt = utils.ReadSecret("salt")

type PasswordHasher interface {
	Hash(password string) string
}

func Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
