package anime

import (
	"context"
	"reflect"
	"testing"

	"github.com/josegpt/iamodb/anime/pb"
)

func TestServer(t *testing.T) {
	repo := InMemoryRepository{
		animes: []Anime{
			{ID: 1, Title: "Tokyo Ghoul", Description: "Testing...", Plot: "Testing..."},
			{ID: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
		},
	}
	server := grpcServer{&repo}

	t.Run("GetAnimes(limit = 2, offset = 0) returns []Anime{{...}, {...}}", func(t *testing.T) {
		req := &pb.GetAnimesRequest{Limit: 2, Offset: 0}
		resp, err := server.GetAnimes(context.Background(), req)

		if err != nil {
			t.Fatalf("unable to process request %v", err)
		}

		want := &pb.GetAnimesResponse{
			Animes: []*pb.Anime{
				&pb.Anime{Id: 1, Title: "Tokyo Ghoul", Description: "Testing...", Plot: "Testing..."},
				&pb.Anime{Id: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
			},
		}

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v want %v", resp, want)
		}
	})

	t.Run("GetAnimes(limit = 2, offset = 0) returns []Anime{{...}, {...}}", func(t *testing.T) {
		req := &pb.GetAnimesRequest{Limit: 2, Offset: 1}
		resp, err := server.GetAnimes(context.Background(), req)

		if err != nil {
			t.Fatalf("unable to process request %v", err)
		}

		want := &pb.GetAnimesResponse{
			Animes: []*pb.Anime{
				&pb.Anime{Id: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
			},
		}

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v want %v", resp, want)
		}
	})

	t.Run("GetAnime(id = 1) returns Anime{...}", func(t *testing.T) {
		req := &pb.GetAnimeRequest{Id: 1}
		resp, err := server.GetAnime(context.Background(), req)

		if err != nil {
			t.Fatalf("unable to process request %v", err)
		}

		want := &pb.GetAnimeResponse{
			Anime: &pb.Anime{
				Id:          1,
				Title:       "Tokyo Ghoul",
				Description: "Testing...",
				Plot:        "Testing...",
			},
		}

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v want %v", resp, want)
		}
	})
}
