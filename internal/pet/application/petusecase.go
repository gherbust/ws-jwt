package application

import "github.com/gherbust/ws-jwt/internal/pet/domain"

type PetUsecaseRepository interface {
	GetPets() (*[]domain.Pet, error)
}

type PetUsecase struct{}

func NewPetUsecase() PetUsecaseRepository {
	return &PetUsecase{}
}

func (p *PetUsecase) GetPets() (*[]domain.Pet, error) {
	pets := []domain.Pet{}
	pets = append(pets, domain.Pet{ID: 1, Name: "Firulais", Especie: "Perro"})
	pets = append(pets, domain.Pet{ID: 2, Name: "Garfield", Especie: "Gato"})
	pets = append(pets, domain.Pet{ID: 3, Name: "Piolin", Especie: "Perro"})
	pets = append(pets, domain.Pet{ID: 4, Name: "Tom", Especie: "Gato"})
	pets = append(pets, domain.Pet{ID: 5, Name: "Lassie", Especie: "Perro"})
	pets = append(pets, domain.Pet{ID: 6, Name: "Snoopy", Especie: "Perro"})
	pets = append(pets, domain.Pet{ID: 7, Name: "Mickey", Especie: "Raton"})
	pets = append(pets, domain.Pet{ID: 8, Name: "Pluto", Especie: "Perro"})
	pets = append(pets, domain.Pet{ID: 9, Name: "Minnie", Especie: "Raton"})
	pets = append(pets, domain.Pet{ID: 10, Name: "Pato Donald", Especie: "Pato"})

	return &pets, nil
}
