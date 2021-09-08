package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RiskyFeryansyahP/bibit-gateway/internal/model"
	"github.com/gorilla/mux"
)

// MovieHandler ...
type MovieHandler struct {
	MovieUC model.UsecaseMovie
}

// NewMovieHandler ...
func NewMovieHandler(s *mux.Router, movieUC model.UsecaseMovie) {
	handler := &MovieHandler{
		MovieUC: movieUC,
	}

	s.HandleFunc("/movies/search", handler.SearchMovie)
	s.HandleFunc("/movie/{id}", handler.DetailMovie)
}

// SearchMovie ...
func (m *MovieHandler) SearchMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	ctx := context.Background()

	v := r.URL.Query()

	keyword := v.Get("s")
	page := v.Get("page")

	input := &model.RequestSearchMovies{
		Searchword: keyword,
		Pagination: page,
	}

	movies, err := m.MovieUC.Search(ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(movies)
}

// DetailMovie ...
func (m *MovieHandler) DetailMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	ctx := context.Background()

	v := mux.Vars(r)

	imdbID := v["id"]

	input := &model.RequestMovieDetail{
		ImdbID: imdbID,
	}

	movie, err := m.MovieUC.DetailMovie(ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(movie)
}
