package main

import "fmt"
func commandPokedex (cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for k,_ := range cfg.pokedex {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}