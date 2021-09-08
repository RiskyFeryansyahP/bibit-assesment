package model

import "context"

// Movie ...
type Movie struct{}

// UsecaseMovie ...
type UsecaseMovie interface {
	Search(ctx context.Context, input *RequestSearchMovies) (*ResponseSearchMovies, error)
	DetailMovie(ctx context.Context, input *RequestMovieDetail) (*MovieRPC, error)
}

// RepositoryGRPCMovie ...
type RepositoryGRPCMovie interface {
	Search(ctx context.Context, input *RequestSearchMovies) (*ResponseSearchMovies, error)
	GetByID(ctx context.Context, input *RequestMovieDetail) (*MovieRPC, error)
}
