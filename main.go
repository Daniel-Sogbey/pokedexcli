package main

import (
	"time"

	"github.com/Daniel-Sogbey/pokedexcli/internal/pokecache"
)

var newPokeCache = pokecache.NewCache(time.Minute * 5)
var pokedex map[string]Pokemon

func main() {
	pokedex = map[string]Pokemon{}
	cfg := &config{
		BaseURL: "https://pokeapi.co/api/v2/location-area/",
	}

	startRepl(cfg)
}
