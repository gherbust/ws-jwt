package application

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gherbust/ws-jwt/internal/auth/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationRepository interface {
	GenerateToken(user domain.User, clientKey string) string
	IsTokenValid(tokenString string) (claim *domain.Claim, isValid bool)
	AuthMiddleware() gin.HandlerFunc
}

type Authentication struct {
	signingKey []byte
	clients    []string
}

func NewAuthentication(signingKey []byte, clients []string) AuthenticationRepository {
	return &Authentication{
		signingKey: signingKey,
		clients:    clients,
	}
}

func (a *Authentication) GenerateToken(user domain.User, clientKey string) string {
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
	tokenString, err := token.SignedString(a.signingKey)
	if err != nil {
		fmt.Printf("%v %v", tokenString, err)
		return ""
	}

	return tokenString
}

func (a *Authentication) IsTokenValid(tokenString string) (claim *domain.Claim, isValid bool) {
	if tokenString == "" {
		fmt.Println("error token string empty")
		return nil, false
	}

	token, err := jwt.ParseWithClaims(tokenString, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return a.signingKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	if claims, ok := token.Claims.(*domain.Claim); ok && token.Valid {
		if !strings.Contains(strings.Join(a.clients, ","), claims.User.ClientKey) {
			fmt.Println("error invalid client key")
			return nil, false
		}
		return claims, token.Valid
	} else {
		fmt.Println(err)
		return nil, false
	}
}

func (a *Authentication) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return a.signingKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(*domain.Claim); ok && token.Valid {
			if !strings.Contains(strings.Join(a.clients, ","), claims.User.ClientKey) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			}

			c.Set("user", claims.User)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		}
	}
}
