package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	BaseURL  string
	Previous string
	Next     string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]

		options := slices.Delete(words, 0, 1)

		if exists {
			err := command.callback(cfg, options...)
			if err != nil {
				fmt.Println(err)
				commandHelp(cfg)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			commandHelp(cfg)
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore",
			description: "See a list of all the Pokémon located there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: " It takes the name of a Pokemon as an argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "It takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"pokedex": {
			name:        "pokedex",
			description: "It takes no arguments but prints a list of all the names of the Pokemon the user has caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
