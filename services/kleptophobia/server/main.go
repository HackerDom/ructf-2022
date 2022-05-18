package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"kleptophobia/models"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	models.UnimplementedKleptophobiaServer
	dbApi *DBApi
}

func (s *server) Register(ctx context.Context, in *models.RegisterRequest) (*models.RegisterReply, error) {
	if err := s.dbApi.register(in.Person, in.Password); err != nil {
		msg := "can not register user: " + err.Error()
		return &models.RegisterReply{
			Status:  models.RegisterReply_FAIL,
			Message: &msg,
		}, nil
	}
	return &models.RegisterReply{Status: models.RegisterReply_OK}, nil
}

func (s *server) GetPublicInfo(ctx context.Context, in *models.GetPublicInfoRequest) (*models.GetPublicInfoReply, error) {
	person, err := s.dbApi.getPublicInfo(in.Username)
	if err != nil {
		msg := "can not get public info: " + err.Error()
		return &models.GetPublicInfoReply{
			Status:  models.GetPublicInfoReply_FAIL,
			Message: &msg,
			Person:  nil,
		}, nil
	}
	return &models.GetPublicInfoReply{
		Status: models.GetPublicInfoReply_OK,
		Person: models.PersonRecordToPublic(person),
	}, nil
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
	models.RegisterKleptophobiaServer(s, &server{dbApi: &dbApi})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
