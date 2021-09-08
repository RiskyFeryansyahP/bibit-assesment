package config

import "os"

// MapConfig ...
type MapConfig struct {
	OmdbURL    string
	OmdbAPIKey string
	Database   *Database
}

// NewMapConfig ...
func NewMapConfig() *MapConfig {
	omdbURL := os.Getenv("OMDB_URL")
	omdbAPIKey := os.Getenv("OMDB_API_KEY")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	database := &Database{
		Host:    dbHost,
		Port:    dbPort,
		User:    dbUser,
		Pass:    dbPass,
		Name:    dbName,
		SSLMode: sslMode,
	}

	return &MapConfig{
		OmdbURL:    omdbURL,
		OmdbAPIKey: omdbAPIKey,
		Database:   database,
	}
}
