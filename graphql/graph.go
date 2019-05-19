package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/josegpt/iamodb/anime"
)

type Server struct {
	animeClient *anime.Client
}

func NewGraphQLServer(animeURL string) (*Server, error) {
	animeClient, err := anime.NewClient(animeURL)
	if err != nil {
		animeClient.Close()
		return nil, err
	}

	return &Server{
		animeClient,
	}, nil
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{s}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
