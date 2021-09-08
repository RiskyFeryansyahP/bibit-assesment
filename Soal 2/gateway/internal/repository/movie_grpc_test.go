package repository

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

	movieServiceMock := mock.NewMockMovieServiceClient(ctrl)

	t.Run("search should be success", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "Batman",
			Pagination: "2",
		}

		movieServiceMock.EXPECT().SearchMovie(ctx, input).Return(&model.ResponseSearchMovies{}, nil).Times(1)

		repo := NewMovieGRPC(movieServiceMock)

		movies, err := repo.Search(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, movies)
	})

	t.Run("search failed can't call movie rpc", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "Batman",
			Pagination: "2",
		}

		movieServiceMock.EXPECT().SearchMovie(ctx, input).Return(nil, fmt.Errorf("failed connect rpc")).Times(1)

		repo := NewMovieGRPC(movieServiceMock)

		_, err := repo.Search(ctx, input)

		require.Error(t, err)
	})

	t.Run("detail movie should be success", func(t *testing.T) {
		input := &model.RequestMovieDetail{
			ImdbID: "tt4853102",
		}

		movieServiceMock.EXPECT().GetMovieDetail(ctx, input).Return(&model.MovieRPC{}, nil).Times(1)

		repo := NewMovieGRPC(movieServiceMock)

		movie, err := repo.GetByID(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, movie)
	})

	t.Run("detail movie failed connect movie rpc", func(t *testing.T) {
		input := &model.RequestMovieDetail{
			ImdbID: "tt4853102",
		}

		movieServiceMock.EXPECT().GetMovieDetail(ctx, input).Return(nil, fmt.Errorf("failed connect rpc")).Times(1)

		repo := NewMovieGRPC(movieServiceMock)

		_, err := repo.GetByID(ctx, input)

		require.Error(t, err)
	})
}
