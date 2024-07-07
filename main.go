package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte{}

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}
	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}
	return nil
}

func main() {

}
func (u *UserClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, u)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error in create token when sining token : %w", err)
	}
	return tokenStr, nil
}
