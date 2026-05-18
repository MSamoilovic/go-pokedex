package main

import (
	"errors"
	"fmt"
	"math/rand"
	pokeApi "go-pokedex/src/internal/pokeApi"
	pokeCache "go-pokedex/src/internal/pokeCache"
	"os"
)


type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args []string) error
}

type config struct {
	Next     string
	Prev     string
	Cache    *pokeCache.Cache
	Pokedex  map[string]pokeApi.Pokemon
	SavePath string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Prints locations in Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous location areas to explore",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"release": {
			name:        "release",
			description: "Release a caught Pokemon",
			callback:    commandRelease,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func commandExit(config *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, args []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%v:\t%v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *config, args []string) error {
	resp, err := pokeApi.GetLocationAreas(config.Next, config.Cache)
	if err != nil {
		return err
	}

	config.Next = resp.Next
	config.Prev = resp.Previous

	for _, location := range resp.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func commandMapb(config *config, args []string) error {
	if config.Prev == "" {
		return errors.New("we're already pulled over, we can't pull over previous any further")
	}

	resp, err := pokeApi.GetLocationAreas(config.Prev, config.Cache)
	if err != nil {
		return err
	}

	config.Next = resp.Next
	config.Prev = resp.Previous

	for _, location := range resp.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func commandCatch(config *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage: catch <pokemon>")
	}

	name := args[0]
	pokemon, err := pokeApi.GetPokemon(name, config.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", colorize(colorYellow, "Throwing a Pokeball at "+name+"..."))

	if rand.Intn(pokemon.BaseExperience+1) < 50 {
		fmt.Printf("%s\n", colorize(colorGreen, colorBold+name+" was caught!"))
		config.Pokedex[name] = pokemon
		if err := savePokedex(config.SavePath, config.Pokedex); err != nil {
			fmt.Printf("Warning: could not save Pokedex: %v\n", err)
		}
	} else {
		fmt.Printf("%s\n", colorize(colorRed, name+" escaped!"))
	}

	return nil
}

func commandRelease(config *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage: release <pokemon>")
	}

	name := args[0]
	if _, ok := config.Pokedex[name]; !ok {
		fmt.Printf("you don't have %s in your Pokedex\n", name)
		return nil
	}

	delete(config.Pokedex, name)
	if err := savePokedex(config.SavePath, config.Pokedex); err != nil {
		fmt.Printf("Warning: could not save Pokedex: %v\n", err)
	}
	fmt.Printf("%s was released!\n", name)

	return nil
}

func commandInspect(config *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage: inspect <pokemon>")
	}

	pokemon, ok := config.Pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", colorize(colorBold, pokemon.Name))
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %s\n", s.Stat.Name, colorize(colorYellow, fmt.Sprintf("%d", s.BaseStat)))
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", typeColor(t.Type.Name))
	}

	return nil
}

func commandPokedex(config *config, args []string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range config.Pokedex {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}

func commandExplore(config *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage: explore <location-area>")
	}

	resp, err := pokeApi.GetLocationArea(args[0], config.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", colorize(colorCyan, args[0]))
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", colorize(colorGreen, encounter.Pokemon.Name))
	}

	return nil
}