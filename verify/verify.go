package verify

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func Verify(secret, token string) error {
	// Parse, validate, verify the signature and return the parsed token
	_, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			// Will receive the parsed token and should return the cryptographic
			// key for verifying the signature.
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	return nil
}
