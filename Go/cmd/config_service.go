package main

import (
	pb "cmd/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

type server struct {
	pb.UnimplementedConfigServiceServer
}

func (s *server) AddConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfig not implemented")
}

func (s *server) GetConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}

func (s *server) GetAllVersionsOfConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Configs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVersionsOfConfig not implemented")
}

func (s *server) GetAllConfigs(ctx context.Context, in *emptypb.Empty) (*pb.Configs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConfigs not implemented")
}

func (s *server) UpdateConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}

func (s *server) DeleteConfig(ctx context.Context, in *pb.ConfigNameRequest) (*pb.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}

func (s *server) UseConfig(in *pb.ConfigNameRequest, stream pb.ConfigService_UseConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method UseConfig not implemented")
}

func (s *server) StopConfigUseForAll(ctx context.Context, in *pb.ConfigNameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopConfigUseForAll not implemented")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed ot listen: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterConfigServiceServer(serv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err = serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
