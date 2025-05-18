package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"godoallserver/internal/service"
)

func RegisterGraphQL(mux &http.ServeMux) {
	srv := handler.NewDefaultServer(service.NewExecutableSchema())

	mux.Handle("/graphql", srv)
	mux.Handle("/playground", playground.Handler("GraphQL playground", "/grpahql"))
}