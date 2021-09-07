package model

import context "context"

// Movie ...
type Movie struct {
	Title      string `json:"Title,omitempty"`
	Year       string `json:"Year,omitempty"`
	ImdbID     string `json:"imdbID,omitempty"`
	Type       string `json:"Type,omitempty"`
	Poster     string `json:"Poster,omitempty"`
	Rated      string `json:"Rated,omitempty"`
	Released   string `json:"Released,omitempty"`
	Runtime    string `json:"Runtime,omitempty"`
	Genre      string `json:"Genre,omitempty"`
	Director   string `json:"Director,omitempty"`
	Writer     string `json:"Writer,omitempty"`
	Actors     string `json:"Actors,omitempty"`
	Plot       string `json:"Plot,omitempty"`
	Language   string `json:"Language,omitempty"`
	Country    string `json:"Country,omitempty"`
	Awards     string `json:"Awards,omitempty"`
	ImdbRating string `json:"imdbRating,omitempty"`
}

// MovieSearch ...
type MovieSearch struct {
	Search []*Movie `json:"Search"`
}

// UsecaseMovie ...
type UsecaseMovie interface {
	SearchValidate(ctx context.Context, keyword, page string) ([]*MovieRPC, error)
	DetailMovie(ctx context.Context, id string) (*MovieRPC, error)
}

// RepositoryMovie ...
type RepositoryMovie interface {
	Search(ctx context.Context, keyword, page string) (*MovieSearch, error)
	GetByID(ctx context.Context, id string) (*Movie, error)
}
