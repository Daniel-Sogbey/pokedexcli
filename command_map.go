package main

import (
	"fmt"
)

func commandMap(cfg *config, options ...string) error {
	if cfg.Next == "" {
		cfg.Next = cfg.BaseURL
	}

	locationArea, err := requester[LocationArea](cfg.Next, "GET", nil)
	if err != nil {
		return err
	}

	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous

	locationAreaBytes, err := toBytes(locationArea)
	if err != nil {
		return err
	}
	newPokeCache.Add(cfg.Next, locationAreaBytes)

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(cfg *config, options ...string) error {
	if cfg.Previous == "" {
		fmt.Println("you’re on the first page")
		return nil
	}

	locationArea, err := getLocationArea(cfg)
	if err != nil {
		return err
	}

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func getLocationArea(cfg *config) (LocationArea, error) {
	var err error
	var locationArea LocationArea

	if val, ok := newPokeCache.Get(cfg.Previous); ok {
		locationArea, err = fromBytes[LocationArea](val)
		if err != nil {
			return LocationArea{}, err
		}

	} else {
		locationArea, err = requester[LocationArea](cfg.Previous, "GET", nil)
		if err != nil {
			return LocationArea{}, err
		}

		cfg.Next = locationArea.Next
		cfg.Previous = locationArea.Previous

		locationAreaBytes, err := toBytes(locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		newPokeCache.Add(cfg.Next, locationAreaBytes)

	}

	return locationArea, nil
}
