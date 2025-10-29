package main

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"example.com/server/poc/server/demo"
)

type userServiceServer struct {
	demo.UnimplementedUserServiceServer
}

func NewUserService() *userServiceServer {
	return &userServiceServer{}
}

func (s *userServiceServer) GetUser(ctx context.Context, _ *emptypb.Empty) (*demo.Person, error) {
	return &demo.Person{
		Name: "Test User",
		Age:  42,
		Status: &demo.NewNewStatus{ // matches server/server.proto
			Status: "OK",
		},
	}, nil
}
