package routes

import (
	"database/sql"

	"Graphql_and_gRPC/app/graph"
	"Graphql_and_gRPC/app/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Db *sql.DB
}

// Defining the Graphql handler
func (_self Route) graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func (_self Route) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (_self Route) RegisterGraphServer() *gin.Engine {
	ginServer := gin.Default()

	ginServer.GET("/", _self.playgroundHandler())
	ginServer.POST("/query", _self.graphqlHandler())

	return ginServer
}
