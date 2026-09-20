package pokeapi

import (
	"net/http"
	"time"

	"github.com/Sandro-GG/Pokedex/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

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
