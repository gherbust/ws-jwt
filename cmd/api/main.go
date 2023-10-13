package main

import (
	"net/http"

	authDomain "github.com/gherbust/ws-jwt/internal/auth/application"
	authInfrastructure "github.com/gherbust/ws-jwt/internal/auth/infrastructure"
	"github.com/gherbust/ws-jwt/internal/pet/application"
	"github.com/gherbust/ws-jwt/internal/pet/infrastructure"
	"github.com/gin-gonic/gin"
)

/*
	func buildRouter() *gin.Engine {
		r := gin.New()
		r.POST("/api/v1/login", dep.AuthHandler.Login)

		client_api := r.Group("/api/v1")
		client_api.Use(dep.AuthRepo.AuthMiddleware())

}
*/
func main() {
	r := gin.Default()
	clients := []string{"cliente 1", "cliente 2", "cliente 3"}
	secret := []byte("No veas es secreto")
	authUseCase := authDomain.NewAuthentication(secret, clients)
	authAndler := authInfrastructure.NewAuthenticatorHandler(authUseCase)

	petUseCase := application.NewPetUsecase()
	petsHandler := infrastructure.NewPetHandler(petUseCase)
	r.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusAccepted, "pong") })
	r.POST("/api/v1/login", authAndler.Login)
	r.GET("/api/v1/valid", authAndler.ValidToken)

	authRouts := r.Group("/api/v1")
	authRouts.Use(authUseCase.AuthMiddleware())
	authRouts.GET("/pets", petsHandler.GetPets)

	r.Run(":8080")
}
