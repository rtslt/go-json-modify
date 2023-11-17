package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/signal"
	"syscall"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

func GetPokemon(pokemon string) []byte {
	// Try to get data from API and manipulate it
	var url = "https://pokeapi.co/api/v2/pokemon?limit=1&offset=0"
	var r, _ = http.Get(url)
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)
	var pokemonResponse map[string]interface{}
	err := json.Unmarshal([]byte(body), &pokemonResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return []byte("Error unmarshaling JSON")
	}
	fmt.Println(pokemonResponse)
	pokemonResponse["ModData"] = "This is a modification"
	fmt.Println(pokemonResponse)

	body, _ = json.Marshal(pokemonResponse)
	return body
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/pokemon/{pokemon}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(GetPokemon(chi.URLParam(r, "pokemon"))))
		// GetPokemon(chi.URLParam(r, "pokemon"))
	})

	fmt.Println("Server running on port 3000")
	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", "3000"),
	}

	go func() {
		fmt.Println(server.ListenAndServe())
	}()

	// Gracefully Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server %s", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}
