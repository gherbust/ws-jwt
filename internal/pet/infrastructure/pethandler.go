package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gherbust/ws-jwt/internal/pet/application"
	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	PetUsecase application.PetUsecaseRepository
}

func NewPetHandler(petUsecase application.PetUsecaseRepository) *PetHandler {
	return &PetHandler{
		PetUsecase: petUsecase,
	}
}

func (p *PetHandler) GetPets(c *gin.Context) {
	pets, err := p.PetUsecase.GetPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user := c.Keys["user"]
	fmt.Println(user)
	c.JSON(http.StatusOK, pets)
}
