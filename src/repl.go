package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokeApi "go-pokedex/src/internal/pokeApi"
	pokeCache "go-pokedex/src/internal/pokeCache"
)


func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		Cache:   pokeCache.NewCache(5 * time.Minute),
		Pokedex: map[string]pokeApi.Pokemon{},
	}

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := scanner.Text()

		if len(words) == 0 {
			continue
		}

		res := cleanInput(words)

		cmd, ok := getCommands()[res[0]]
		if !ok {
			fmt.Printf("Unrecognized command :%v\n", words)
			continue
		}

		err := cmd.callback(&config, res[1:])
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Split(strings.TrimSpace(text), " ")
	clean_words := []string{}

	for _, word := range words {
		if strings.Contains(word, " ") || word == "" {
			continue
		}

		clean_words = append(clean_words, strings.ToLower(word))
	}

	return clean_words
}