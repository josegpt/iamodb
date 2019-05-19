//go:generate protoc ./anime.proto --go_out=plugins=grpc:./pb
package anime

import (
	"context"
	"fmt"
	"net"

	"github.com/josegpt/iamodb/anime/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAnimeServiceServer(server, &grpcServer{s})
	reflection.Register(server)
	return server.Serve(listen)
}

func (s *grpcServer) GetAnimes(ctx context.Context, r *pb.GetAnimesRequest) (*pb.GetAnimesResponse, error) {
	res, err := s.service.GetAnimes(ctx, r.Limit, r.Offset)
	if err != nil {
		return nil, err
	}
	var animes []*pb.Anime
	for _, a := range res {
		animes = append(
			animes,
			&pb.Anime{
				Id:          a.ID,
				Title:       a.Title,
				Description: a.Description,
				Plot:        a.Plot,
			},
		)
	}

	return &pb.GetAnimesResponse{Animes: animes}, nil
}

func (s *grpcServer) GetAnime(ctx context.Context, r *pb.GetAnimeRequest) (*pb.GetAnimeResponse, error) {
	a, err := s.service.GetAnime(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAnimeResponse{
		Anime: &pb.Anime{
			Id:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			Plot:        a.Plot,
		},
	}, nil
}
