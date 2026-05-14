package main

import (
	"encoding/json"
	"os"

	pokeApi "go-pokedex/src/internal/pokeApi"
)

func loadPokedex(path string) map[string]pokeApi.Pokemon {
	data, err := os.ReadFile(path)
	if err != nil {
		return map[string]pokeApi.Pokemon{}
	}

	pokedex := map[string]pokeApi.Pokemon{}
	if err := json.Unmarshal(data, &pokedex); err != nil {
		return map[string]pokeApi.Pokemon{}
	}

	return pokedex
}

func savePokedex(path string, pokedex map[string]pokeApi.Pokemon) error {
	data, err := json.MarshalIndent(pokedex, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
