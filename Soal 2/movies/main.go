package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/RiskyFeryansyahP/bibit-movies/config"
	"github.com/RiskyFeryansyahP/bibit-movies/ent/migrate"
	"github.com/RiskyFeryansyahP/bibit-movies/internal/handler"
	"github.com/RiskyFeryansyahP/bibit-movies/internal/repository"
	"github.com/RiskyFeryansyahP/bibit-movies/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	port = fmt.Sprintf(":%s", port)

	ctx := context.Background()

	cfg := config.NewMapConfig()

	client, err := cfg.ConnectDB()
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen movie service: %v", err)
	}

	err = client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s := grpc.NewServer()

	repo := repository.NewMovieRepository(cfg, client)
	uc := usecase.NewMovieUsecase(repo)
	handler.NewMovieHandler(s, uc)

	log.Println("serve movie service at :8081")

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve gRPC: %s", err)
	}
}
