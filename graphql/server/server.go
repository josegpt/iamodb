//go:generate go run github.com/99designs/gqlgen -v
package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/josegpt/iamodb/graphql"
)

type Config struct {
	AnimeURL string
}

const defaultPort = 8080

func main() {
	cfg := Config{"localhost:5000"}
	// err := envconfig.Process("", &cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	s, err := graphql.NewGraphQLServer(cfg.AnimeURL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
