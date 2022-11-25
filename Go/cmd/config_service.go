package main

import (
	"cmd/internal/db"
	"cmd/internal/repository"
	"cmd/internal/server"
	service2 "cmd/internal/service"
	pb "cmd/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

func main() {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed ot listen: %v", err)
	}

	repo := repository.NewConfigRepository(client)
	service := service2.NewDistributedConfigService(repo)
	serv := grpc.NewServer()
	pb.RegisterConfigServiceServer(serv, server.NewServer(service))
	log.Printf("server listening at %v", lis.Addr())
	if err = serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
