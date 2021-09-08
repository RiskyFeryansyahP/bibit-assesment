package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/RiskyFeryansyahP/bibit-movies/config"
	"github.com/RiskyFeryansyahP/bibit-movies/ent"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/stretchr/testify/require"
)

func TestMovieRepository(t *testing.T) {
	ctx := context.Background()

	client, err := sqliteConnection()
	require.NoError(t, err)

	err = client.Schema.Create(ctx)
	require.NoError(t, err)

	os.Setenv("OMDB_URL", "http://www.omdbapi.com/")
	os.Setenv("OMDB_API_KEY", "faf7e5bb")

	cfg := config.NewMapConfig()

	repo := NewMovieRepository(cfg, client)

	t.Run("Search should be success", func(t *testing.T) {
		keyword := "Batman"
		page := "1"

		movies, err := repo.Search(ctx, keyword, page)

		require.NoError(t, err)
		require.NotNil(t, movies)
	})

	t.Run("Search should failed request timeout", func(t *testing.T) {
		// trying to make timeout operation in database and will return error
		ctxTimeout, cancel := context.WithTimeout(ctx, 0*time.Millisecond)
		defer cancel()

		keyword := "Batman"
		page := "1"

		movies, err := repo.Search(ctxTimeout, keyword, page)

		require.Error(t, err)
		require.Nil(t, movies)
	})

	t.Run("Detail movie should be success", func(t *testing.T) {
		imdbID := "tt4853102"

		movie, err := repo.GetByID(ctx, imdbID)

		require.NoError(t, err)
		require.NotNil(t, movie)
		require.Equal(t, "Batman: The Killing Joke", movie.Title)
	})
}

func sqliteConnection() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	return client, nil
}
