package main

import "fmt"

func commandHelp(cfg *config, options ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
