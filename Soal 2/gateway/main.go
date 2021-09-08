package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RiskyFeryansyahP/bibit-gateway/internal/handler"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/model"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/repository"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/usecase"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	movieServicePort := os.Getenv("MOVIE_SERVICE")
	movieHost := os.Getenv("MOVIE_SERVICE_HOST")

	if port == "" {
		port = "8080"
	}

	if movieServicePort == "" {
		movieServicePort = "8081"
	}

	if movieHost == "" {
		movieHost = "localhost"
	}

	port = fmt.Sprintf(":%s", port)
	movieServicePort = fmt.Sprintf("%s:%s", movieHost, movieServicePort)

	conn, err := grpc.Dial(movieServicePort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect grpc movies server: %s", err)
	}
	defer conn.Close()

	m := model.NewMovieServiceClient(conn)

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	movieRepo := repository.NewMovieGRPC(m)
	movieUC := usecase.NewMovieUsecase(movieRepo)
	handler.NewMovieHandler(s, movieUC)

	log.Println("server gateway running at localhost:8080")

	http.ListenAndServe(port, r)
}
