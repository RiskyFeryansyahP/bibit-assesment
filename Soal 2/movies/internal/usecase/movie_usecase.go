package usecase

import (
	"context"
	"fmt"

	"github.com/RiskyFeryansyahP/bibit-movies/internal/model"
)

// MovieUsecase ...
type MovieUsecase struct {
	MovieRepo model.RepositoryMovie
}

// NewMovieUsecase ...
func NewMovieUsecase(movieRepo model.RepositoryMovie) model.UsecaseMovie {
	return &MovieUsecase{
		MovieRepo: movieRepo,
	}
}

// SearchMovie ...
func (mu *MovieUsecase) SearchMovie(ctx context.Context, keyword, page string) ([]*model.Movie, error) {
	if keyword == "" {
		return nil, fmt.Errorf("keyword can't be empty")
	}

	if page == "" {
		page = "1"
	}

	movies, err := mu.MovieRepo.Search(ctx, keyword, page)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
