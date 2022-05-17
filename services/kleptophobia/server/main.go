package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "kleptophobia/models"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedKleptophobiaServer
	dbApi *DBApi
}

func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	if err := s.dbApi.register(in.Person); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.RegisterReply{Status: pb.RegisterReply_OK}, nil
}

func main() {
	dbApi := DBApi{}
	dbApi.init("localhost", 5432, "myusername", "mypassword", "myusername")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterKleptophobiaServer(s, &server{dbApi: &dbApi})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
