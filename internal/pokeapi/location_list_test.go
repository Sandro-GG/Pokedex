package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestListLocationAreas(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"next": null, "previous": null, "results": [{"name": "canalave-city-area"}]}`))
	}))
	defer server.Close()

	c := NewClient(time.Second)
	url := server.URL
	result, err := c.ListLocationAreas(&url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result.Results))
	}
	if result.Results[0].Name != "canalave-city-area" {
		t.Errorf("expected canalave-city-area, got %s", result.Results[0].Name)
	}
}

func TestListLocationAreaPagination(t *testing.T) {
	nextUrl := "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"next": "` + nextUrl + `", "previous": null, "results": [{"name": "mt-coronet-1f-route-216"}]}`))
	}))
	defer server.Close()

	c := NewClient(time.Second)
	url := server.URL
	result, err := c.ListLocationAreas(&url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Next == nil {
		t.Fatalf("expected Next to be non-nil, got nil")
	}

	if *result.Next != nextUrl {
		t.Fatalf("expected %s, got %s", nextUrl, *result.Next)
	}

	if result.Previous != nil {
		t.Fatalf("expected Previous to be nil, got %s", *result.Previous)
	}
}
