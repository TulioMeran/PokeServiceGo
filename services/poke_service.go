package services

import "TulioMeran/example/models"

type PokeService interface {
	GetAll() ([]models.Pokemon, error)
	GetPokemonByName(name string) (models.Pokemon, error)
}
