package main

import (
	"log"
	"net"

	"github.com/RiskyFeryansyahP/bibit-movies/internal/handler"
	"github.com/RiskyFeryansyahP/bibit-movies/internal/repository"
	"github.com/RiskyFeryansyahP/bibit-movies/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen movie service: %v", err)
	}

	s := grpc.NewServer()

	repo := repository.NewMovieRepository()
	uc := usecase.NewMovieUsecase(repo)
	handler.NewMovieHandler(s, uc)

	log.Println("serve movie service at :8081")

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve gRPC: %s", err)
	}
}
