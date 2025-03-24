package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch (cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	pokemonName := args[0]

	// Pokemon Request
	pokemonResp, err := cfg.pokeapiClient.Pokemon(pokemonName)
	if err != nil {
		return nil
	}

	res :=rand.Intn(pokemonResp.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if (40 < res) {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	} 

	
	fmt.Printf("%s was caught!\n", pokemonResp.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.pokedex[pokemonResp.Name] = pokemonResp
	return nil
}