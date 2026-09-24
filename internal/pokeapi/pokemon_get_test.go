package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestPokemonGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"name": "pikachu",
			"height": 4,
			"weight": 60
		}`))
	}))
	defer server.Close()

	c := NewClient(time.Second)
	c.pokemonUrl = server.URL + "/"

	result, err := c.PokemonGet("pikachu")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "pikachu" {
		t.Errorf("expected pikachu, got %s", result.Name)
	}
	if result.Height != 4 {
		t.Errorf("expected height 4, got %d", result.Height)
	}
	if result.Weight != 60 {
		t.Errorf("expected weight 60, got %d", result.Weight)
	}
}

func TestPokemonGetNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	c := NewClient(time.Second)
	c.pokemonUrl = server.URL + "/"

	_, err := c.PokemonGet("pikapika")
	if err == nil {
		t.Fatalf("expected an error for 404, got nil")
	}

	expected := "pokemon name not found"

	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}

}
