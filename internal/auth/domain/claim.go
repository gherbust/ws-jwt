package domain

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	User `json:"user"`
	jwt.RegisteredClaims
}
