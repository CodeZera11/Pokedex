package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please enter a pokemon name")
	}

	name := args[0]

	pokemon, err := c.pokeapiClient.GetPokemon(name)

	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if chance > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	c.pokedex[name] = pokemon

	return nil
}