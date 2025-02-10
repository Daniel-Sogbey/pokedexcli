package main

import "fmt"

func commandPokedex(cfg *config, options ...string) error {

	if len(pokedex) == 0 {
		fmt.Println("you have not caught any pokedex yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pk := range pokedex {
		fmt.Println("- ", pk.Name)
	}
	return nil
}
