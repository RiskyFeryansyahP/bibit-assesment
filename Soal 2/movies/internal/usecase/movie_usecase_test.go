package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/RiskyFeryansyahP/bibit-movies/internal/mock"
	"github.com/RiskyFeryansyahP/bibit-movies/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestSearchValidate(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	repo := mock.NewMockRepositoryMovie(ctrl)

	t.Run("search validate should be success", func(t *testing.T) {
		keyword := "Batman"
		page := "1"

		result := &model.MovieSearch{
			Search: []*model.Movie{
				{Title: "Batman: The Killing Joke", ImdbID: "tt4853102", Year: "2016", Type: "movie", Poster: "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},
			},
		}

		repo.EXPECT().Search(ctx, keyword, page).Return(result, nil).Times(1)

		usecase := NewMovieUsecase(repo)

		movies, err := usecase.SearchValidate(ctx, keyword, page)

		require.NoError(t, err)
		require.NotNil(t, movies)
	})

	t.Run("search validate should be success with empty page", func(t *testing.T) {
		keyword := "Batman"
		page := "1"

		repo.EXPECT().Search(ctx, keyword, page).Return(&model.MovieSearch{}, nil).Times(1)

		usecase := NewMovieUsecase(repo)

		movies, err := usecase.SearchValidate(ctx, keyword, "")

		require.NoError(t, err)
		require.NotNil(t, movies)
	})

	t.Run("search validate failed empty keyword", func(t *testing.T) {
		keyword := ""
		page := "1"

		usecase := NewMovieUsecase(repo)

		_, err := usecase.SearchValidate(ctx, keyword, page)

		require.Error(t, err)
		require.Equal(t, "keyword can't be empty", err.Error())
	})

	t.Run("search validate failed from repository", func(t *testing.T) {
		keyword := "Batman"
		page := "1"

		repo.EXPECT().Search(ctx, keyword, page).Return(nil, fmt.Errorf("request time out")).Times(1)

		usecase := NewMovieUsecase(repo)

		_, err := usecase.SearchValidate(ctx, keyword, page)

		require.Error(t, err)
	})

	t.Run("detail movie should be success", func(t *testing.T) {
		imdbID := "tt4853102"

		repo.EXPECT().GetByID(ctx, imdbID).Return(&model.Movie{}, nil).Times(1)

		usecase := NewMovieUsecase(repo)

		movie, err := usecase.DetailMovie(ctx, imdbID)

		require.NoError(t, err)
		require.NotNil(t, movie)
	})

	t.Run("detail movie failed empty id", func(t *testing.T) {
		imdbID := ""

		usecase := NewMovieUsecase(repo)

		_, err := usecase.DetailMovie(ctx, imdbID)

		require.Error(t, err)
		require.Equal(t, "id can't be empty", err.Error())
	})

	t.Run("detail movie failed from repository", func(t *testing.T) {
		imdbID := "tt4853102"

		repo.EXPECT().GetByID(ctx, imdbID).Return(nil, fmt.Errorf("request time out")).Times(1)

		usecase := NewMovieUsecase(repo)

		_, err := usecase.DetailMovie(ctx, imdbID)

		require.Error(t, err)
	})
}
