package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"example.com/server/poc/server/demo"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register service implementation
	demo.RegisterUserServiceServer(s, NewUserService())

	log.Printf("gRPC server listening on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
