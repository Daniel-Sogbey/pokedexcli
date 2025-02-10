package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, options ...string) error {
	if len(options) == 0 {
		return errors.New("add a pokemon location area")
	}
	name := options[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := requester[Pokemon](url, "GET", nil)
	if err != nil {
		return err
	}

	catchRate := 100 - (pokemon.BaseExperience / 100)

	randomNum := rand.Intn(101)
	if randomNum < catchRate {
		pokedex[name] = pokemon
		fmt.Printf("%s was caught\n", name)
	}

	return nil
}
