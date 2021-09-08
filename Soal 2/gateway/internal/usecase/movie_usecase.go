package usecase

import (
	"context"
	"fmt"

	"github.com/RiskyFeryansyahP/bibit-gateway/internal/model"
)

// MovieUsecase ..
type MovieUsecase struct {
	MovieRepo model.RepositoryGRPCMovie
}

// NewMovieUsecase ...
func NewMovieUsecase(movieRepo model.RepositoryGRPCMovie) model.UsecaseMovie {
	return &MovieUsecase{
		MovieRepo: movieRepo,
	}
}

// Search ...
func (m *MovieUsecase) Search(ctx context.Context, input *model.RequestSearchMovies) (*model.ResponseSearchMovies, error) {
	if input.Searchword == "" {
		return nil, fmt.Errorf("search word can't be empty")
	}

	if input.Pagination == "" {
		input.Pagination = "1"
	}

	movies, err := m.MovieRepo.Search(ctx, input)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// DetailMovie ...
func (m *MovieUsecase) DetailMovie(ctx context.Context, input *model.RequestMovieDetail) (*model.MovieRPC, error) {
	if input.ImdbID == "" {
		return nil, fmt.Errorf("ImdbID can't be empty")
	}

	movie, err := m.MovieRepo.GetByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return movie, nil
}
