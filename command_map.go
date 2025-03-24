package main

import (
	"fmt"
)



func commandMap (cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return nil
	}
	
	cfg.nextLocationsUrl = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB (cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil || *cfg.prevLocationURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return nil
	}
	
	cfg.nextLocationsUrl = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}