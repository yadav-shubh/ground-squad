package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/yadav-shubh/base-middleware/graph/generated"
	"github.com/yadav-shubh/base-middleware/graph/resolvers"
	"github.com/yadav-shubh/base-middleware/utils"
	"go.uber.org/zap"
	"log"
	"net/http"
)

const defaultPort = "0.0.0.0:8080"

func main() {
	aadr := defaultPort

	resolver := &resolvers.Resolver{}
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})

	http.Handle("/middleware", playground.Handler("GraphQL playground", "/middleware/query"))
	http.Handle("/middleware/query", srv)

	log.Printf("connect for GraphQL playground at %s", aadr)
	err := http.ListenAndServe(aadr, nil)
	if err != nil {
		utils.Log.Error("failed to connect for GraphQL playground", zap.Error(err))
	}
}
