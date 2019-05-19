package anime

import "context"

type Service interface {
	GetAnimes(ctx context.Context, limit, offset uint64) ([]Anime, error)
	GetAnime(ctx context.Context, id uint64) (*Anime, error)
}

type Anime struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Plot        string `json:"plot"`
}

type animeService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &animeService{r}
}

func (s *animeService) GetAnimes(ctx context.Context, limit, offset uint64) ([]Anime, error) {
	return s.repository.GetAnimes(ctx, limit, offset)
}

func (s *animeService) GetAnime(ctx context.Context, id uint64) (*Anime, error) {
	return s.repository.GetAnime(ctx, id)
}
