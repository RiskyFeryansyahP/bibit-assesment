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

// SearchValidate ...
func (mu *MovieUsecase) SearchValidate(ctx context.Context, keyword, page string) ([]*model.MovieRPC, error) {
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

	movieRPCs := make([]*model.MovieRPC, len(movies.Search))

	for k, v := range movies.Search {
		movieRPC := &model.MovieRPC{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}

		movieRPCs[k] = movieRPC
	}

	return movieRPCs, nil
}
