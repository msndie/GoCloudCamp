package server

import (
	"cmd/internal/service"
	pb "cmd/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedConfigServiceServer
	service *service.DistributedConfigService
}

func NewServer(service *service.DistributedConfigService) *Server {
	return &Server{service: service}
}

func (s *Server) AddConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfig not implemented")
}

func (s *Server) GetConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}

func (s *Server) GetAllVersionsOfConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Configs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVersionsOfConfig not implemented")
}

func (s *Server) GetAllConfigs(ctx context.Context, in *emptypb.Empty) (*pb.Configs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConfigs not implemented")
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
