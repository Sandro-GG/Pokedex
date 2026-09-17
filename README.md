# Pokedex CLI

A command-line Pokedex built in Go, using the [PokeAPI](https://pokeapi.co/).

## Features

- Interactive REPL for exploring Pokemon data
- `map` / `mapb` commands to paginate through location areas
- In-memory caching layer to reduce redundant API requests

## Caching

Requests to the PokeAPI are cached in memory using a custom `pokecache` package.
- Cache entries expire automatically after a configurable interval
- Thread-safe for concurrent access (read/write locks)
- Reduces network calls when revisiting previously fetched pages

## Installation

```
git clone https://github.com/Sandro-GG/Pokedex
cd Pokedex
go build -o pokedex
```

## Usage

```
./pokedex
```

Available commands:
- `map` - list next page of location areas
- `mapb` - list previous page of location areas
- `help` - show available commands
- `exit` - quit the program

## Testing

Run all tests from the repo root:

\`\`\`
go test ./...
\`\`\`