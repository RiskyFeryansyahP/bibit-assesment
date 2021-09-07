package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Movies Service")

	l, err := net.Listen("tcp", "8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve gRPC: %s", err)
	}
}
