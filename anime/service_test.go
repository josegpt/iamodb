package anime

import (
	"context"
	"reflect"
	"testing"
)

type InMemoryRepository struct {
	animes []Anime
}

func (m *InMemoryRepository) Close() {}

func (m *InMemoryRepository) Ping() error {
	return nil
}

func (m *InMemoryRepository) GetAnimes(ctx context.Context, limit, offset uint64) ([]Anime, error) {
	return m.animes[offset:limit], nil
}

func (m *InMemoryRepository) GetAnime(ctx context.Context, id uint64) (*Anime, error) {
	var anime Anime
	for _, a := range m.animes {
		if a.ID == id {
			anime = Anime{
				ID:          a.ID,
				Title:       a.Title,
				Description: a.Description,
				Plot:        a.Plot,
			}
		}
	}
	return &anime, nil
}

func TestService(t *testing.T) {
	repo := InMemoryRepository{
		animes: []Anime{
			{ID: 1, Title: "Tokyo Ghoul", Description: "Testing...", Plot: "Testing..."},
			{ID: 2, Title: "Death Parade", Description: "Testing...", Plot: "Testing..."},
		},
	}
	service := NewService(&repo)

	t.Run("GetAnimes(limit = 2, offset = 0) returns []Anime[{...}, {...}]", func(t *testing.T) {
		got, err := service.GetAnimes(context.Background(), 2, 0)

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

	t.Run("GetAnimes(limit = 2, offset = 1) returns []Anime[{...}, {...}]", func(t *testing.T) {
		got, err := service.GetAnimes(context.Background(), 2, 1)

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

	t.Run("GetAnime(id = 1) returns &Anime{...}", func(t *testing.T) {
		got, err := service.GetAnime(context.Background(), 1)

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
}
