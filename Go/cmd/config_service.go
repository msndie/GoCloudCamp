package main

import (
	"cmd/internal/db"
	"cmd/internal/repository"
	"cmd/internal/server"
	service2 "cmd/internal/service"
	pb "cmd/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	port   = flag.Int("port", 9090, "The server port")
	dbname = flag.String("db_name", "GoCloud", "Name of the database")
	dbURI  = flag.String("db_uri", "mongodb://localhost:27017/", "URI for database connection")
)

func main() {
	flag.Parse()
	client, err := db.GetMongoClient(*dbURI)
	if err != nil {
		log.Fatal(err.Error())
	}

	f := func() {
		log.Print("Shutting server down")
		_ = client.Disconnect(context.TODO())
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		f()
		os.Exit(0)
	}()

	defer f()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed ot listen: %v", err)
	}

	repo := repository.NewConfigRepository(client.Database(*dbname))
	service := service2.NewDistributedConfigService(repo)
	serv := grpc.NewServer()
	pb.RegisterConfigServiceServer(serv, server.NewServer(service))
	log.Printf("server listening at %v", lis.Addr())
	if err = serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
