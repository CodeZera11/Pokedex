package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 { 
		return errors.New("please enter a location to search")
	}

	location := args[0]

	fmt.Printf("Exploring %s...\n", location)

	resp, err := cfg.pokeapiClient.GetLocation(location)

	if err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")
	for _, val := range resp.PokemonEncounters{
		fmt.Println("- ", val.Pokemon.Name)
	}

	return nil
}