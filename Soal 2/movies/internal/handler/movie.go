package handler

import (
	"context"

	"github.com/RiskyFeryansyahP/bibit-movies/internal/model"
	"google.golang.org/grpc"
)

// MovieHandler ...
type MovieHandler struct {
	MovieUC model.UsecaseMovie
}

// NewMovieHandler ...
func NewMovieHandler(server *grpc.Server, movieUC model.UsecaseMovie) {
	handler := &MovieHandler{
		MovieUC: movieUC,
	}

	model.RegisterMovieServiceServer(server, handler)
}

// SearchMovie ...
func (mh *MovieHandler) SearchMovie(ctx context.Context, req *model.RequestSearchMovies) (*model.ResponseSearchMovies, error) {
	keyword := req.Searchword
	page := req.Pagination

	movies, err := mh.MovieUC.SearchValidate(ctx, keyword, page)
	if err != nil {
		return nil, err
	}

	resp := &model.ResponseSearchMovies{
		Movies: movies,
	}

	return resp, nil
}

// GetMovieDetail ...
func (mh *MovieHandler) GetMovieDetail(ctx context.Context, req *model.RequestMovieDetail) (*model.MovieRPC, error) {
	imdbID := req.ImdbID

	movie, err := mh.MovieUC.DetailMovie(ctx, imdbID)
	if err != nil {
		return nil, err
	}

	return movie, nil
}
