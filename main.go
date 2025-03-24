package main

import (
	"time"

	"github.com/robertstipp/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)
	cfg := &config{
		pokedex: make(map[string]pokeapi.RespPokemon),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}

