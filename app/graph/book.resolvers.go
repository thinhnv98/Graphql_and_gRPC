package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"os"

	"Graphql_and_gRPC/app/graph/generated"
	"Graphql_and_gRPC/app/graph/model"
	"Graphql_and_gRPC/app/grpc/pb"
	"google.golang.org/grpc"
)

func (r *mutationResolver) CreateBookAPI(ctx context.Context, input model.NewBookRequest) (*model.BookResponse, error) {
	//Register connection to book-microservices
	bookServerConn, err := grpc.DialContext(ctx, os.Getenv("BOOK_SERVICE_HOST")+":"+os.Getenv("BOOK_SERVICE_PORT"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	bookClient := pb.NewBookServicesClient(bookServerConn)

	//Call book-services
	bookResponse, err := bookClient.CreateBook(ctx, &pb.NewBookRequest{
		Name:  input.Name,
		Total: int32(input.Total),
	})
	log.Print("Send CreateBook request value: ", input)
	if err != nil {
		return nil, err
	}

	return &model.BookResponse{
		ID:    int(bookResponse.GetId()),
		Total: int(bookResponse.GetTotal()),
	}, nil
}

func (r *queryResolver) Book(ctx context.Context, id *int) (*model.BookResponse, error) {
	return &model.BookResponse{
		ID:    1,
		Total: 2,
	}, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.BookResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
