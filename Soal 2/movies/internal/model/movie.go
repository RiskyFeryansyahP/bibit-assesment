package model

import context "context"

// UsecaseMovie ...
type UsecaseMovie interface {
	SearchMovie(ctx context.Context, keyword, page string) ([]*Movie, error)
}

// RepositoryMovie ...
type RepositoryMovie interface {
	Search(ctx context.Context, keyword, page string) ([]*Movie, error)
}
