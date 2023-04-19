package util

import (
	"bytes"
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

	// Get sample pokeapi response.
	pokeapiResponse := getSamplePokeapiResponse(c)

	// Parse pokemon successfully with pokeapiResponse
	parsedPokemon, err := ParsePokemon(pokeapiResponse)
	c.NoError(err)

	// Get expected api response
	expectedApiResponse := getSampleApiResponse(c)

	// Expect parsedPokemon to be equal to expectedPokemonResponse.
	c.Equal(expectedApiResponse, parsedPokemon)
}

func TestParserNotFoundPokemonTypeFailure(t *testing.T) {
	c := require.New(t)

	// Get sample pokeapi response.
	pokeapiResponse := getSamplePokeapiResponse(c)

	// Put PokemonType as empty array.
	pokeapiResponse.PokemonType = []models.PokemonType{}

	// Returns ErrNotFoundPokemonType
	_, err := ParsePokemon(pokeapiResponse)
	c.Error(err, ErrNotFoundPokemonType)
}

func TestParserNotFoundPokemonTypeNameFailure(t *testing.T) {
	c := require.New(t)

	// Get sample pokeapi response.
	pokeapiResponse := getSamplePokeapiResponse(c)

	// Put TypeName as empty string.
	pokeapiResponse.PokemonType[0].RefType.Name = ""

	// Returns ErrNotFoundPokemonType
	_, err := ParsePokemon(pokeapiResponse)
	c.Error(err, ErrNotFoundPokemonTypeName)
}

func getSamplePokeapiResponse(c *require.Assertions) *models.PokeApiPokemonResponse {
	// Get pokeapi response from samples
	fileContent, err := ioutil.ReadFile("samples/pokeapi_response.json")
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

func getSampleApiResponse(c *require.Assertions) *models.Pokemon {
	// Get expectedPokemonResponse from samples
	fileContent, err := ioutil.ReadFile("samples/api_response.json")
	c.NoError(err)

	// Write fileContent to byte buffer successfully
	var buf bytes.Buffer
	_, err = buf.Write(fileContent)
	c.NoError(err)

	// Decode file content from buffer to expectedPokemonResponse model
	var expectedPokemonResponse *models.Pokemon
	decoder := json.NewDecoder(&buf)
	err = decoder.Decode(&expectedPokemonResponse)
	c.NoError(err)

	return expectedPokemonResponse
}
