package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, options ...string) error {
	if len(options) == 0 {
		return errors.New("add a pokemon location area")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + options[0]
	fmt.Printf("Exploring %s...\n", options[0])

	pokemonList, err := requester[PokemonList](url, "GET", nil)
	if err != nil {
		return err
	}

	for _, pl := range pokemonList.PokemonEncounters {
		fmt.Println(pl.Pokemon.Name)
	}
	return nil
}
