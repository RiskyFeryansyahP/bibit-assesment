package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RiskyFeryansyahP/bibit-gateway/internal/mock"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestSearchMovie(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	usecase := mock.NewMockUsecaseMovie(ctrl)

	r := mux.NewRouter()
	subR := r.PathPrefix("/api/v1").Subrouter()

	NewMovieHandler(subR, usecase)

	t.Run("search movie http call should success", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "Batman",
			Pagination: "2",
		}

		usecase.EXPECT().Search(ctx, input).Return(&model.ResponseSearchMovies{}, nil).Times(1)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/movies/search?s=Batman&page=2", nil)
		require.NoError(t, err)

		req.URL.Query().Set("s", input.Searchword)
		req.URL.Query().Set("page", input.Pagination)

		hdlr := &MovieHandler{
			MovieUC: usecase,
		}

		rr := httptest.NewRecorder()

		handler := subR.HandleFunc("/movies/search", hdlr.SearchMovie).Methods(http.MethodGet)

		handler.GetHandler().ServeHTTP(rr, req)

		body := rr.Body.Bytes()

		var resp *model.ResponseSearchMovies

		err = json.Unmarshal(body, &resp)

		require.NoError(t, err)
		require.NotNil(t, resp)
	})

	t.Run("search movie http call should error empty searchword", func(t *testing.T) {
		input := &model.RequestSearchMovies{
			Searchword: "",
			Pagination: "2",
		}

		usecase.EXPECT().Search(ctx, input).Return(nil, fmt.Errorf("search word can't be empty")).Times(1)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/movies/search?page=2", nil)
		require.NoError(t, err)

		req.URL.Query().Set("s", input.Searchword)
		req.URL.Query().Set("page", input.Pagination)

		hdlr := &MovieHandler{
			MovieUC: usecase,
		}

		rr := httptest.NewRecorder()

		handler := subR.HandleFunc("/movies/search", hdlr.SearchMovie).Methods(http.MethodGet)

		handler.GetHandler().ServeHTTP(rr, req)

		body := rr.Body.Bytes()

		var resp *model.ResponseSearchMovies

		err = json.Unmarshal(body, &resp)

		require.NoError(t, err)
		require.Nil(t, resp.Movies)
	})
}

func TestDetailMovie(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()

	usecase := mock.NewMockUsecaseMovie(ctrl)

	r := mux.NewRouter()
	subR := r.PathPrefix("/api/v1").Subrouter()

	NewMovieHandler(subR, usecase)

	t.Run("detail movie http call should be success", func(t *testing.T) {
		input := &model.RequestMovieDetail{
			ImdbID: "tt4853102",
		}

		usecase.EXPECT().DetailMovie(ctx, input).Return(&model.MovieRPC{}, nil).Times(1)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/movie/tt4853102", nil)
		require.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{"id": input.ImdbID})

		req.Header.Set("Content-Type", "application/json")

		hdlr := &MovieHandler{
			MovieUC: usecase,
		}

		rr := httptest.NewRecorder()

		handler := subR.HandleFunc("/movies/{id}", hdlr.DetailMovie).Methods(http.MethodGet)

		handler.GetHandler().ServeHTTP(rr, req)

		body := rr.Body.Bytes()

		var resp *model.MovieRPC

		err = json.Unmarshal(body, &resp)

		require.NoError(t, err)
		require.NotNil(t, resp)
	})

	t.Run("detail movie http call should be failed empty ImdbID", func(t *testing.T) {
		input := &model.RequestMovieDetail{
			ImdbID: "",
		}

		usecase.EXPECT().DetailMovie(ctx, input).Return(nil, fmt.Errorf("ImdbID can't be empty")).Times(1)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/movie", nil)
		require.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{"id": input.ImdbID})

		req.Header.Set("Content-Type", "application/json")

		hdlr := &MovieHandler{
			MovieUC: usecase,
		}

		rr := httptest.NewRecorder()

		handler := subR.HandleFunc("/movies/{id}", hdlr.DetailMovie).Methods(http.MethodGet)

		handler.GetHandler().ServeHTTP(rr, req)

		body := rr.Body.Bytes()

		var resp *model.MovieRPC

		err = json.Unmarshal(body, &resp)

		require.NoError(t, err)
		require.Empty(t, resp)
	})
}
