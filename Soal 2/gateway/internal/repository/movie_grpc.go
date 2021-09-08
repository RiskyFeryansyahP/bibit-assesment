package repository

import (
	"context"

	"github.com/RiskyFeryansyahP/bibit-gateway/internal/model"
)

// MovieGRPC ...
type MovieGRPC struct {
	Movie model.MovieServiceClient
}

// NewMovieGRPC ...
func NewMovieGRPC(movie model.MovieServiceClient) model.RepositoryGRPCMovie {
	return &MovieGRPC{
		Movie: movie,
	}
}

// Search ...
func (m *MovieGRPC) Search(ctx context.Context, input *model.RequestSearchMovies) (*model.ResponseSearchMovies, error) {
	movies, err := m.Movie.SearchMovie(ctx, input)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// GetByID ...
func (m *MovieGRPC) GetByID(ctx context.Context, input *model.RequestMovieDetail) (*model.MovieRPC, error) {
	movie, err := m.Movie.GetMovieDetail(ctx, input)
	if err != nil {
		return nil, err
	}

	return movie, nil
}
