package reverse_proxy

import (
	"context"
	"flag"
	"net/http"
	"os"

	"Graphql_and_gRPC/app/grpc/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type ReverseProxy struct {
}

func (_self ReverseProxy) RegisterProxy(ctx context.Context) (http.Handler, error) {
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	grpcMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterBookServicesHandlerFromEndpoint(ctx, grpcMux, *flag.String("book-grpc-endpoint", ":"+os.Getenv("BOOK_SERVICE_PORT"), "gRPC server endpoint"), opts); err != nil {
		return nil, err
	}

	//if err := pb.RegisterBookServicesHandlerServer(ctx, grpcMux, nil); err != nil{
	//	return nil, err
	//}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	return mux, nil
}
