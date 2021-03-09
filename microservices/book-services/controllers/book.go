package controllers

import (
	"context"

	"Graphql_and_gRPC/microservices/book-services/pb"
)

type BookController struct {
	pb.UnimplementedBookServicesServer
}

func (_self BookController) CreateBook(ctx context.Context, input *pb.NewBookRequest) (*pb.NewBookResponse, error) {
	return &pb.NewBookResponse{
		Id:    100,
		Total: 100,
	}, nil
}
