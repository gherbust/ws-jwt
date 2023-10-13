package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gherbust/ws-jwt/internal/auth/domain"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	user := domain.User{
		Name:      "gerry",
		ClientKey: "cliente 2",
		Role:      "admin",
	}
	token := GenerateToken(user, "No veas es secreto", "cliente 2")

	fmt.Println(token)

	calims, isValid := IsTokenValid(token, "No veas es secreto")
	if isValid {
		fmt.Println(calims)
	}
}

func GenerateToken(user domain.User, signingKey, clientKey string) string {
	claims := domain.Claim{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    clientKey,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		fmt.Printf("%v %v", tokenString, err)
		return ""
	}

	return tokenString
}

func IsTokenValid(tokenString, signingKey string) (claim *domain.Claim, isValid bool) {
	clients := []string{"cliente 2", "cliente 2"}
	if tokenString == "" {
		fmt.Println("error token string empty")
		return nil, false
	}

	token, err := jwt.ParseWithClaims(tokenString, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	if claims, ok := token.Claims.(*domain.Claim); ok && token.Valid {
		if !strings.Contains(strings.Join(clients, ","), claims.User.ClientKey) {
			fmt.Println("error invalid client key")
			return nil, false
		}
		return claims, token.Valid
	} else {
		fmt.Println(err)
		return nil, false
	}
}
