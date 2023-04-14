package util

import (
	"catching-pokemons/models"
	"errors"
)

var (
	// ErrNotFoundPokemonType occurs when the type array in pokeapi response it's not found
	ErrNotFoundPokemonType = errors.New("pokemon type array not found")
	// ErrNotFoundPokemonTypeName occurs when we found type struct but no name
	ErrNotFoundPokemonTypeName = errors.New("pokemon type name not found")
)

func ParsePokemon(apiPokemon models.PokeApiPokemonResponse) (*models.Pokemon, error) {
	// Check we have pokemon
	if len(apiPokemon.PokemonType) < 1 {
		return &models.Pokemon{}, ErrNotFoundPokemonType
	}

	// Check we have pokemon name
	if apiPokemon.PokemonType[0].RefType.Name == "" {
		return &models.Pokemon{}, ErrNotFoundPokemonTypeName
	}

	// Only get desired abilities.
	abilitiesMap := map[string]int{}
	for _, stat := range apiPokemon.Stats {
		parsedAbilityName, ok := models.AllowedAbilities[stat.Stat.Name]
		if !ok {
			continue
		}

		abilitiesMap[parsedAbilityName] = stat.BaseStat
	}

	// Compose our parsed pokemon with important fields.
	parsedPokemon := &models.Pokemon{
		Id:        apiPokemon.Id,
		Name:      apiPokemon.Name,
		Power:     apiPokemon.PokemonType[0].RefType.Name,
		Abilities: abilitiesMap,
	}

	return parsedPokemon, nil
}
