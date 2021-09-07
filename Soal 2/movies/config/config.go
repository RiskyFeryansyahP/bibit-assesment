package config

import "os"

// MapConfig ...
type MapConfig struct {
	OmdbURL    string
	OmdbAPIKey string
}

// NewMapConfig ...
func NewMapConfig() *MapConfig {
	omdbURL := os.Getenv("OMDB_URL")
	omdbAPIKey := os.Getenv("OMDB_API_KEY")

	return &MapConfig{
		OmdbURL:    omdbURL,
		OmdbAPIKey: omdbAPIKey,
	}
}
