package controller

import (
	"catching-pokemons/models"
	"catching-pokemons/util"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const API_TIMEOUT = 5

var (
	// ErrEmptyId occurs when id param is empty on request
	ErrEmptyId = errors.New("id is required")
	// ErrPokeAPIFailure occurs when PokeAPIResponse sends an error to our request.
	ErrPokeAPIFailure = errors.New("pokeapi failure")
	// ErrPokemonNotFound occurs when pokemon wasn't found.
	ErrPokemonNotFound = errors.New("pokemon not found")
	// ErrInternal occurs when something unexpected happened and we want to ofuscate it.
	ErrInternal = errors.New("internal error. Please contact and administrator")
)

// GetPokemon gets pokemon from provided name id.
func GetPokemon(c *gin.Context) {

	// Check Id is passed.
	id := c.Param("id")
	if id == "" {
		respondwithJSON(c, http.StatusBadRequest, ErrEmptyId.Error())
		return
	}

	// Gets the Pokemon from PokeAPI
	apiPokemon, err := getDecodedPokeapiResponse(id, API_TIMEOUT)
	if err != nil {
		respondwithJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Parses pokeapiResponse object to compose desired API response object.
	parsedPokemon, err := util.ParsePokemon(apiPokemon)
	if err != nil {
		respondwithJSON(c, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondwithJSON(c, http.StatusOK, parsedPokemon)
}

func getDecodedPokeapiResponse(id string, timeout int) (*models.PokeApiPokemonResponse, error) {
	// Gets the Pokemon from PokeAPI
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	response, err := client.Get(request)
	if err != nil {
		return &models.PokeApiPokemonResponse{}, ErrInternal
	}

	// If pokemon wasn't found
	if response.StatusCode == http.StatusNotFound {
		return &models.PokeApiPokemonResponse{}, ErrPokemonNotFound
	}

	// If PokeAPI had an internal failure
	if response.StatusCode != http.StatusOK {
		return &models.PokeApiPokemonResponse{}, ErrPokeAPIFailure
	}

	// Reads data from body and decodes fields inside object.
	defer response.Body.Close()
	var apiPokemon *models.PokeApiPokemonResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&apiPokemon)
	if err != nil {
		return &models.PokeApiPokemonResponse{}, ErrInternal
	}

	return apiPokemon, nil
}
