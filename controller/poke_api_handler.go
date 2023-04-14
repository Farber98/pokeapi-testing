package controller

import (
	"catching-pokemons/models"
	"catching-pokemons/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPokemon gets pokemon from provided name id.
func GetPokemon(c *gin.Context) {
	// Gets the Pokemon from PokeAPI
	id := c.Param("id")
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
	response, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}

	// Reads data from body and decodes the desired fields inside API response object.
	defer response.Body.Close()
	var apiPokemon models.PokeApiPokemonResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&apiPokemon)
	if err != nil {
		log.Fatal(err)
	}

	parsedPokemon, err := util.ParsePokemon(apiPokemon)
	if err != nil {
		respondwithJSON(c, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondwithJSON(c, http.StatusOK, parsedPokemon)
}
