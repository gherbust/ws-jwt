package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/gherbust/ws-jwt/internal/auth/application"
	"github.com/gherbust/ws-jwt/internal/auth/domain"
	"github.com/gin-gonic/gin"
)

type AuthenticatorHandler struct {
	Auth application.AuthenticationRepository
}

func NewAuthenticatorHandler(auth application.AuthenticationRepository) AuthenticatorHandler {
	return AuthenticatorHandler{
		Auth: auth,
	}
}

func (a *AuthenticatorHandler) Login(c *gin.Context) {
	var user domain.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Parametros de acceso erroneos")
		return
	}

	if user.Name != "gerry" && user.Password != "Pass123$" {
		c.JSON(http.StatusNotAcceptable, nil)
		return
	}

	user.Password = ""
	user.Role = "Admin"
	alias := "cliente 1"
	token := a.Auth.GenerateToken(user, alias)
	responseToken := domain.ResponseToken{
		Token: token,
	}
	//c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusAccepted, responseToken)
	return
}

func (a *AuthenticatorHandler) ValidToken(c *gin.Context) {
	claims, isvalid := a.Auth.IsTokenValid(c.Request.Header.Get("Authorization"))
	if !isvalid {
		c.JSON(http.StatusUnauthorized, nil)
	}
	c.JSON(http.StatusAccepted, claims.User)
}
