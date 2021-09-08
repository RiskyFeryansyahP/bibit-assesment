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
	movieService := os.Getenv("MOVIE_SERVICE")

	if port == "" {
		port = "8080"
	}

	if movieService == "" {
		movieService = "8081"
	}

	port = fmt.Sprintf(":%s", port)
	movieService = fmt.Sprintf(":%s", movieService)

	conn, err := grpc.Dial(movieService, grpc.WithInsecure())
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
