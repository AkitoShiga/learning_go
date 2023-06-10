package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
)

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()

	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))

	return
}
