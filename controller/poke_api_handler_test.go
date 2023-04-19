package controller

import (
	"bytes"
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDecodedPokeApiResponseSuccess(t *testing.T) {
	c := require.New(t)

	// Get pokemon from pokeapi
	pokemon, err := getDecodedPokeapiResponse("ditto")
	c.NoError(err)

	// Get expected api response
	expectedResponse := getSamplePokeapiResponse(c)

	// Expect parsedPokemon to be equal to expectedPokemonResponse.
	c.Equal(expectedResponse, pokemon)
}
func TestGetDecodedPokeApiResponsePokemonNotFoundFailure(t *testing.T) {
	c := require.New(t)

	// Get inexsistent pokemon from pokeapi
	_, err := getDecodedPokeapiResponse("not a pokemon")
	c.Error(err, ErrPokemonNotFound)
}

func getSamplePokeapiResponse(c *require.Assertions) *models.PokeApiPokemonResponse {
	// Get pokeapi response from samples
	fileContent, err := ioutil.ReadFile("../util/samples/pokeapi_response.json")
	c.NoError(err)

	// Write fileContent to byte buffer successfully
	var buf bytes.Buffer
	_, err = buf.Write(fileContent)
	c.NoError(err)

	// Decode file content from buffer to PokeApiPokemonResponse model
	var apiPokemon *models.PokeApiPokemonResponse
	decoder := json.NewDecoder(&buf)
	err = decoder.Decode(&apiPokemon)
	c.NoError(err)

	return apiPokemon
}
