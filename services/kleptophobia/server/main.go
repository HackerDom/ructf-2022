package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"kleptophobia/models"
	"kleptophobia/utils"
	"log"
	"net"
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

func (s *server) GetPublicInfo(ctx context.Context, in *models.GetByUsernameRequest) (*models.GetPublicInfoReply, error) {
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

func (s *server) GetEncryptedFullInfo(ctx context.Context, in *models.GetByUsernameRequest) (*models.GetEncryptedFullInfoReply, error) {
	encryptedFullInfo, err := s.dbApi.getEncryptedFullInfo(in.Username)
	if err != nil {
		msg := "can not get public info: " + err.Error()
		return &models.GetEncryptedFullInfoReply{
			Status:            models.GetEncryptedFullInfoReply_FAIL,
			Message:           &msg,
			EncryptedFullInfo: nil,
		}, nil
	}
	return &models.GetEncryptedFullInfoReply{
		Status:            models.GetEncryptedFullInfoReply_OK,
		Message:           nil,
		EncryptedFullInfo: encryptedFullInfo,
	}, nil
}

func (s *server) Ping(ctx context.Context, in *models.PingBody) (*models.PingBody, error) {
	return in, nil
}

func main() {
	configFilename := flag.String("config", "dev_config.json", "server config")
	flag.Parse()

	var serverConfig models.ServerConfig
	utils.InitConfig[*models.ServerConfig](*configFilename, &serverConfig)

	dbApi := DBApi{}
	dbApi.init(serverConfig.PgConfig)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverConfig.GrpcPort))
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