package controller

import (
	"bytes"
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetDecodedPokeApiResponseSuccess(t *testing.T) {
	c := require.New(t)

	// Get pokemon from pokeapi
	pokemon, err := getDecodedPokeapiResponse("ditto", API_TIMEOUT)
	c.NoError(err)

	// Get expected api response
	expectedResponse := getSamplePokeapiResponse(c)

	// Expect parsedPokemon to be equal to expectedPokemonResponse.
	c.Equal(expectedResponse, pokemon)
}

func TestGetDecodedPokeApiResponseSuccessMocked(t *testing.T) {
	c := require.New(t)

	// Activate mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Get pokeapi response from samples
	fileContent, err := ioutil.ReadFile("../util/samples/pokeapi_response.json")
	c.NoError(err)

	// Set request that is going to be made.
	id := "ditto"
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	// Mock the call.
	mockCall(200, "GET", request, string(fileContent))

	// Call the method with /id and get the decoded response.
	pokemon, err := getDecodedPokeapiResponse(id, API_TIMEOUT)
	c.NoError(err)

	// Get decoded pokemon from file.
	expectedPokemon := decodePokemonFromFileContent(c, fileContent)

	// Assert that mocked pokemon is equal to sample pokemon.
	c.Equal(pokemon, expectedPokemon)
}
func TestGetDecodedPokeApiResponsePokemonNotFoundFailure(t *testing.T) {
	c := require.New(t)

	// Get inexsistent pokemon from pokeapi
	_, err := getDecodedPokeapiResponse("not a pokemon", API_TIMEOUT)
	c.EqualError(err, ErrPokemonNotFound.Error())
}

func TestGetDecodedPokeApiResponsePokemonNotFoundFailureMocked(t *testing.T) {
	c := require.New(t)

	// Activate mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Set request that is going to be called.
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", "cualquierita")

	// Mock the call.
	mockCall(http.StatusNotFound, "GET", request, "")

	// Get inexsistent pokemon from pokeapi
	_, err := getDecodedPokeapiResponse("cualquierita", API_TIMEOUT)
	c.EqualError(err, ErrPokemonNotFound.Error())
}

func TestGetDecodedPokeApiResponseInternalServerErrorDecodingJsonMocked(t *testing.T) {
	c := require.New(t)

	// Activate mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Set request that is going to be called.
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", "ditto")

	// Mock the call.
	mockCall(http.StatusOK, "GET", request, "invalid json")

	// Get inexsistent pokemon from pokeapi
	_, err := getDecodedPokeapiResponse("ditto", API_TIMEOUT)
	c.EqualError(err, ErrInternal.Error())
}

/* func TestGetDecodedPokeApiResponseInternalServerErrorHttpClient(t *testing.T) {
	c := require.New(t)

	// Get inexsistent pokemon from pokeapi
	_, err := getDecodedPokeapiResponse("ditto", 0)
	c.EqualError(err, ErrInternal.Error())
} */

func TestGetDecodedPokeApiResponsePokeapiFailureMocked(t *testing.T) {
	c := require.New(t)

	// Activate mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Set request that is going to be made.
	id := "ditto"
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	// Mock the call.
	mockCall(http.StatusInternalServerError, "GET", request, "invalid json")

	// Get inexsistent pokemon from pokeapi
	pokemon, err := getDecodedPokeapiResponse(id, API_TIMEOUT)
	fmt.Println(pokemon)
	c.EqualError(err, ErrPokeAPIFailure.Error())
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

func decodePokemonFromFileContent(c *require.Assertions, fileContent []byte) *models.PokeApiPokemonResponse {
	// Write fileContent to byte buffer successfully
	var buf bytes.Buffer
	_, err := buf.Write(fileContent)
	c.NoError(err)

	// Decode file content from buffer to PokeApiPokemonResponse model
	var apiPokemon *models.PokeApiPokemonResponse
	decoder := json.NewDecoder(&buf)
	err = decoder.Decode(&apiPokemon)
	c.NoError(err)

	return apiPokemon
}

func mockCall(status int, method, request, response string) {

	// Set sample as response when /id is called.
	httpmock.RegisterResponder(method, request, httpmock.NewStringResponder(status, response))
}
