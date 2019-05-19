package anime

import (
	"context"
	"reflect"
	"testing"

	"github.com/josegpt/iamodb/anime/pb"
	"google.golang.org/grpc"
)

type MockAnimeServer struct {
	animes []*pb.Anime
}

func (m *MockAnimeServer) GetAnimes(ctx context.Context, in *pb.GetAnimesRequest, opts ...grpc.CallOption) (*pb.GetAnimesResponse, error) {
	return &pb.GetAnimesResponse{Animes: m.animes[in.Offset:in.Limit]}, nil
}

func (m *MockAnimeServer) GetAnime(ctx context.Context, in *pb.GetAnimeRequest, opts ...grpc.CallOption) (*pb.GetAnimeResponse, error) {
	var anime pb.Anime
	for _, a := range m.animes {
		if a.Id == in.Id {
			anime = pb.Anime{
				Id:          a.Id,
				Title:       a.Title,
				Description: a.Description,
				Plot:        a.Plot,
			}
		}
	}
	return &pb.GetAnimeResponse{Anime: &anime}, nil
}

func TestClient(t *testing.T) {
	server := &MockAnimeServer{
		animes: []*pb.Anime{
			&pb.Anime{Id: 1, Title: "Tokyo Ghoul", Description: "Testing...", Plot: "Testing..."},
			&pb.Anime{Id: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
		},
	}
	client := Client{
		nil,
		server,
	}

	t.Run("GetAnimes(limit = 2, offset = 0) returns []Anime[{...}, {...}]", func(t *testing.T) {
		got, err := client.GetAnimes(context.Background(), 2, 0)

		if err != nil {
			t.Fatalf("unable to process your request %v", err)
		}

		want := []Anime{
			{ID: 1, Title: "Tokyo Ghoul", Description: "Testing...", Plot: "Testing..."},
			{ID: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("GetAnimes(limit = 2, offset = 1) returns []Anime[{...}]", func(t *testing.T) {
		got, err := client.GetAnimes(context.Background(), 2, 1)

		if err != nil {
			t.Fatalf("unable to process your request %v", err)
		}

		want := []Anime{
			{ID: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("GetAnime(id = 1) &returns Anime[{...}]", func(t *testing.T) {
		got, err := client.GetAnime(context.Background(), 1)

		if err != nil {
			t.Fatalf("unable to process your request %v", err)
		}

		want := &Anime{
			ID:          1,
			Title:       "Tokyo Ghoul",
			Description: "Testing...",
			Plot:        "Testing...",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("GetAnime(id = 2) returns &Anime[{...}]", func(t *testing.T) {
		got, err := client.GetAnime(context.Background(), 2)

		if err != nil {
			t.Fatalf("unable to process your request %v", err)
		}

		want := &Anime{
			ID:          2,
			Title:       "Death Parade",
			Description: "Testing...",
			Plot:        "Testing...",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
