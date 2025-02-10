package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, options ...string) error {
	if len(options) == 0 {
		return errors.New("add a pokemon location area")
	}
	name := options[0]

	pk, ok := pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	}

	fmt.Println("Name: ", pk.Name)
	fmt.Println("Height: ", pk.Height)
	fmt.Println("Weight: ", pk.Weight)
	fmt.Println("Stats: ")

	for _, st := range pk.Stats {
		fmt.Printf("- %s: %v\n", st.Stat.Name, st.BaseStat)
	}

	fmt.Println("Types: ")

	for _, tp := range pk.Types {
		fmt.Println("- ", tp.Type.Name)
	}

	return nil
}
