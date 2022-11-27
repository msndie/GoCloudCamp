package server

import (
	"cmd/internal/service"
	pb "cmd/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type Server struct {
	pb.UnimplementedConfigServiceServer
	service *service.DistributedConfigService
}

func NewServer(service *service.DistributedConfigService) *Server {
	return &Server{service: service}
}

func (s *Server) AddConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	b, err := s.service.AddConfig(in)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	if b {
		return in, nil
	}
	return nil, status.Errorf(codes.AlreadyExists, "Config already exists")
}

func (s *Server) GetConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Config, error) {
	c, err := s.service.FindConfig(in.GetService())
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	if c != nil {
		return c, nil
	}
	return nil, status.Errorf(codes.NotFound, "Config not found")
}

func (s *Server) GetAllVersionsOfConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Configs, error) {
	configs, err := s.service.GetAllVersionsOfConfig(in.Service)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	if configs != nil {
		return &pb.Configs{Configs: configs}, nil
	}
	return nil, status.Errorf(codes.NotFound, "Config not found")
}

func (s *Server) GetAllConfigs(ctx context.Context, in *emptypb.Empty) (*pb.Configs, error) {
	configs, err := s.service.GetAllConfigs()
	if err != nil {
		log.Print(err.Error())
		return nil, status.Errorf(codes.Internal, "Error occurred")
	}
	return &pb.Configs{Configs: configs}, nil
}

func (s *Server) UpdateConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}

func (s *Server) DeleteConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}

func (s *Server) UseConfig(in *pb.ConfigNameRequest, stream pb.ConfigService_UseConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method UseConfig not implemented")
}

func (s *Server) StopConfigUseForAll(ctx context.Context, in *pb.ConfigNameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopConfigUseForAll not implemented")
}
