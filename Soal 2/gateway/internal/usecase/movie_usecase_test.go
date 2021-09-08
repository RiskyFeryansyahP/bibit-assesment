package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/RiskyFeryansyahP/bibit-gateway/internal/mock"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	repo := mock.NewMockRepositoryGRPCMovie(ctrl)

	t.Run("search usecase should success", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "Batman",
			Pagination: "2",
		}

		repo.EXPECT().Search(ctx, input).Return(&model.ResponseSearchMovies{}, nil).Times(1)

		usecase := NewMovieUsecase(repo)

		movies, err := usecase.Search(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, movies)
	})

	t.Run("search usecase should success with emtpy pagination", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "Batman",
			Pagination: "",
		}

		repo.EXPECT().Search(ctx, input).Return(&model.ResponseSearchMovies{}, nil).Times(1)

		usecase := NewMovieUsecase(repo)

		movies, err := usecase.Search(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, movies)
	})

	t.Run("search usecase should failed because emtpy searchword", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "",
			Pagination: "2",
		}

		usecase := NewMovieUsecase(repo)

		_, err := usecase.Search(ctx, input)

		require.Error(t, err)
	})

	t.Run("search usecase should failed from repository", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "Batman",
			Pagination: "2",
		}

		repo.EXPECT().Search(ctx, input).Return(nil, fmt.Errorf("can't connect movie rpc")).Times(1)

		usecase := NewMovieUsecase(repo)

		_, err := usecase.Search(ctx, input)

		require.Error(t, err)
	})
}

func TestDetailMovie(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	repo := mock.NewMockRepositoryGRPCMovie(ctrl)

	t.Run("detail movie call should be success", func(t *testing.T) {
		input := &model.RequestMovieDetail{
			ImdbID: "tt4853102",
		}

		repo.EXPECT().GetByID(ctx, input).Return(&model.MovieRPC{}, nil).Times(1)

		usecase := NewMovieUsecase(repo)

		movie, err := usecase.DetailMovie(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, movie)
	})

	t.Run("detail movie call should be failed empty imdbID", func(t *testing.T) {
		input := &model.RequestMovieDetail{
			ImdbID: "",
		}

		usecase := NewMovieUsecase(repo)

		_, err := usecase.DetailMovie(ctx, input)

		require.Error(t, err)
	})

	t.Run("detail movie call should be failed error rpc from repository", func(t *testing.T) {
		input := &model.RequestMovieDetail{
			ImdbID: "tt4853102",
		}

		repo.EXPECT().GetByID(ctx, input).Return(nil, fmt.Errorf("can't connect movie rpc")).Times(1)

		usecase := NewMovieUsecase(repo)

		_, err := usecase.DetailMovie(ctx, input)

		require.Error(t, err)
	})
}
