package anime

import (
	"context"

	"github.com/josegpt/iamodb/anime/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	services pb.AnimeServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewAnimeServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) GetAnimes(ctx context.Context, limit, offset uint64) ([]Anime, error) {
	r, err := c.services.GetAnimes(ctx, &pb.GetAnimesRequest{Limit: limit, Offset: offset})

	if err != nil {
		return nil, err
	}

	var animes []Anime
	for _, a := range r.Animes {
		animes = append(animes, Anime{
			ID:          a.Id,
			Title:       a.Title,
			Description: a.Description,
			Plot:        a.Plot,
		})
	}

	return animes, nil
}

func (c *Client) GetAnime(ctx context.Context, id uint64) (*Anime, error) {
	r, err := c.services.GetAnime(ctx, &pb.GetAnimeRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return &Anime{
		ID:          r.Anime.Id,
		Title:       r.Anime.Title,
		Description: r.Anime.Description,
		Plot:        r.Anime.Plot,
	}, nil
}
