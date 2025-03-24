package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/robertstipp/pokedexcli/internal/pokeapi"
)

type config struct {
	pokedex map[string]pokeapi.RespPokemon
	pokeapiClient pokeapi.Client
	nextLocationsUrl *string
	prevLocationURL *string
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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg,args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name: "catch <pokemon_name>",
			description: "Catch pokemon",
			callback: commandCatch,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Map explore",
			callback: commandExplore,
		},
		"pokedex": {
			name: "pokedex",
			description: "View pokedex",
			callback: commandPokedex,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "Inspect pokemon",
			callback: commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Map display",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Map back",
			callback: commandMapB,
		},
		
	}
}

