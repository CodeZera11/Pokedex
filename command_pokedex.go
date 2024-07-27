package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, args ...string) error {

	pokedex := c.pokedex

	if len(pokedex) == 0 {
		return errors.New("you have not caught any pokemons")
	}

	fmt.Println("Your pokedex")

	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}