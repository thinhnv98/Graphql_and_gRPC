package main

import (
	reverse_proxy "Graphql_and_gRPC/app/reverse-proxy"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"Graphql_and_gRPC/app/config"
	"Graphql_and_gRPC/app/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Database
	configDB := config.Database{}
	db := configDB.InitDB()

	// Routes
	router := routes.Route{Db: db}

	// Startup Server
	go RunMicroservicesGateway()
	router.RegisterGraphServer().Run()
}

func RunMicroservicesGateway() {
	// Instance Server
	server := gin.Default()

	//gRPC Gate-way
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	reverseProxy := reverse_proxy.ReverseProxy{}
	gateway, err := reverseProxy.RegisterProxy(ctx)
	if err != nil {
		panic(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	server.Any("/*ProxyPath", gin.WrapH(gateway))
	server.Run(":" + os.Getenv("GATEWAY_PORT"))
}
