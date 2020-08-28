package handlers

import (
    "github.com/99designs/gqlgen/handler"
    "github.com/kc1491/go-gql-server/internal/gql"
    "github.com/kc1491/go-gql-server/internal/gql/resolvers"
    "github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler() gin.HandleFunc {
	// New ExecutableSchema and Config are in the generated.go file

	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}

	h := handler.GraphQL(gql.NewExecutableSchena(c))

	return func(C* gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL server.", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, C.Request)
	}
}
