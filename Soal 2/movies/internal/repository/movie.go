package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RiskyFeryansyahP/bibit-movies/internal/model"
)

// MovieRepository ...
type MovieRepository struct{}

// NewMovieRepository ...
func NewMovieRepository() model.RepositoryMovie {
	return &MovieRepository{}
}

// Search ...
func (mr *MovieRepository) Search(ctx context.Context, keyword, page string) (*model.MovieSearch, error) {
	var movies *model.MovieSearch

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&s=%s&page=%s", keyword, page)

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
