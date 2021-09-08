package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RiskyFeryansyahP/bibit-gateway/internal/handler"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/model"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/repository"
	"github.com/RiskyFeryansyahP/bibit-gateway/internal/usecase"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	wait := time.Second * 15

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

	addr := fmt.Sprintf("0.0.0.0%s", port)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c // block until we receive signal

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutdown api gateway")
	os.Exit(0)
}
