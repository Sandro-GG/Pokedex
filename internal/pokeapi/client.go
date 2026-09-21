package pokeapi

import (
	"net/http"
	"time"

	"github.com/Sandro-GG/Pokedex/internal/pokecache"
)

const baseLocationUrl = "https://pokeapi.co/api/v2/location-area/"
const basePokemonUrl = "https://pokeapi.co/api/v2/pokemon/"

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(interval),
	}
}
