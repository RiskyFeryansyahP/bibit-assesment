package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RiskyFeryansyahP/bibit-movies/config"
	"github.com/RiskyFeryansyahP/bibit-movies/ent"
	"github.com/RiskyFeryansyahP/bibit-movies/internal/model"

	_ "github.com/joho/godotenv/autoload" // autoload env
)

// MovieRepository ...
type MovieRepository struct {
	Config *config.MapConfig
	Client *ent.Client
}

// NewMovieRepository ...
func NewMovieRepository(cfg *config.MapConfig, client *ent.Client) model.RepositoryMovie {
	return &MovieRepository{
		Config: cfg,
		Client: client,
	}
}

// Search ...
func (mr *MovieRepository) Search(ctx context.Context, keyword, page string) (*model.MovieSearch, error) {
	var movies *model.MovieSearch

	err := mr.Client.LogSearch.Create().
		SetKeyword(keyword).
		SetPage(page).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?apikey=%s&s=%s&page=%s", mr.Config.OmdbURL, mr.Config.OmdbAPIKey, keyword, page)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(b, &movies)

	return movies, nil
}

// GetByID ...
func (mr *MovieRepository) GetByID(ctx context.Context, id string) (*model.Movie, error) {
	var movie *model.Movie

	// get detail movie using imdbID
	url := fmt.Sprintf("%s?apikey=%s&i=%s", mr.Config.OmdbURL, mr.Config.OmdbAPIKey, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(b, &movie)

	return movie, nil
}
