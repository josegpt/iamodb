package graphql

import (
	"context"
	"fmt"
	"strconv"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Animes(ctx context.Context, pagination *PaginationInput) ([]*Anime, error) {
	limit, offset := uint64(pagination.Limit), uint64(pagination.Offset)
	animeList, err := r.server.animeClient.GetAnimes(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	var animes []*Anime
	for _, a := range animeList {
		anime := Anime{
			ID:          fmt.Sprintf("%d", a.ID),
			Title:       a.Title,
			Description: a.Description,
			Plot:        a.Plot,
		}
		animes = append(animes, &anime)
	}

	return animes, nil
}

func (r *queryResolver) Anime(ctx context.Context, id string) (*Anime, error) {
	uId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return nil, err
	}

	anime, err := r.server.animeClient.GetAnime(ctx, uId)
	if err != nil {
		return nil, err
	}
	return &Anime{
		ID:          fmt.Sprintf("%d", anime.ID),
		Title:       anime.Title,
		Description: anime.Description,
		Plot:        anime.Plot,
	}, nil
}
