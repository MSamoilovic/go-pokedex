# Go Pokedex

A project built as part of the [Boot.dev](https://boot.dev) Go course.

A CLI-based Pokedex that communicates with the [PokeAPI](https://pokeapi.co) to let you explore the Pokemon world, catch Pokemon, and inspect your collection. Responses are cached in memory to avoid redundant API calls.

## Commands

| Command | Description |
|---|---|
| `map` | Display the next 20 location areas |
| `mapb` | Display the previous 20 location areas |
| `explore <area>` | List all Pokemon found in a location area |
| `catch <pokemon>` | Attempt to catch a Pokemon |
| `inspect <pokemon>` | Show details about a caught Pokemon |
| `pokedex` | List all caught Pokemon |
| `help` | Show available commands |
| `exit` | Exit the program |

## How to run

```bash
go run ./src/
```
