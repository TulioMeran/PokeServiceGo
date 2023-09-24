package services

import (
	"TulioMeran/example/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeServiceImpl struct{}

func (p PokeServiceImpl) GetAll() ([]models.Pokemon, error) {
	var apiResponse models.Response
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		fmt.Errorf("Error happnend: %v", err.Error())
		return nil, err
	}

	json.NewDecoder(resp.Body).Decode(&apiResponse)

	return apiResponse.Results, nil

}

func (p PokeServiceImpl) GetPokemonByName(name string) (models.Pokemon, error) {
	var pokemon models.Pokemon
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", name)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Error happnend: %v", err.Error())
		return pokemon, err
	}

	err = json.NewDecoder(resp.Body).Decode(&pokemon)

	if err != nil {
		fmt.Errorf("Error happnend: %v", err.Error())
		return pokemon, err
	}

	return pokemon, nil

}
