package main

import (
	"time"

	"github.com/codezera11/pokedex/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Second)
	cfg := &config{
		pokedex: make(map[string]pokeapi.RespPokemon),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}

