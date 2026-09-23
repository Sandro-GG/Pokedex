# Pokedex CLI

A command-line Pokedex built in Go, using the [PokeAPI](https://pokeapi.co/).

## Features

- Interactive REPL for exploring the Pokemon world
- Pagination through Pokemon world location areas
- Area exploration to see which Pokemon reside there
- Catching mechanics with experience-based catch rates
- Persistent in-memory Pokedex to track caught Pokemon
- Custom in-memory caching layer to minimize redundant API requests

## Caching

Requests to the PokeAPI are cached in memory using a custom `pokecache` package.
- Cache entries expire automatically after a configurable interval
- Thread-safe for concurrent access (read/write locks)
- Reduces network calls when revisiting previously fetched pages

## Installation

```bash
git clone git@github.com:Sandro-GG/Pokedex.git
cd Pokedex
go build -o pokedex
```

## Usage

```bash
./pokedex
```

Available commands:
- `map` - Display the next 20 location areas
- `mapb` - Display the previous 20 location areas
- `explore <area_name>` - List all Pokemon found in a given area
- `catch <pokemon_name> [ball_type]` - Attempt to catch a Pokemon using a specific ball (`pokeball`, `greatball`, `ultraball`, `masterball`; defaults to `pokeball`)
- `inspect <pokemon_name>` - View details (height, weight, stats, types) of a Pokemon you've caught
- `pokedex` - Show all the Pokemon you've caught
- `help` - Show available commands
- `exit` - Quit the program


## Testing

Run all tests from the repo root:

```bash
go test ./...
```