package main

import (
	"context"
	"flag"
	"fmt"
	"kleptophobia/models"
	"kleptophobia/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	models.UnimplementedKleptophobiaServer
	dbApi *DBApi
}

func (s *server) Register(ctx context.Context, in *models.RegisterReq) (*models.RegisterRsp, error) {
	if err := s.dbApi.register(in.Person, in.Password); err != nil {
		msg := "can not register user: " + err.Error()
		return &models.RegisterRsp{
			Status:  models.RegisterRsp_FAIL,
			Message: &msg,
		}, nil
	}
	return &models.RegisterRsp{Status: models.RegisterRsp_OK}, nil
}

func (s *server) GetPublicInfo(ctx context.Context, in *models.GetByUsernameReq) (*models.GetPublicInfoRsp, error) {
	person, err := s.dbApi.getPublicInfo(in.Username)
	if err != nil {
		msg := "can not get public info: " + err.Error()
		return &models.GetPublicInfoRsp{
			Status:  models.GetPublicInfoRsp_FAIL,
			Message: &msg,
			Person:  nil,
		}, nil
	}
	return &models.GetPublicInfoRsp{
		Status: models.GetPublicInfoRsp_OK,
		Person: models.PersonRecordToPublic(person),
	}, nil
}

func (s *server) GetEncryptedFullInfo(ctx context.Context, in *models.GetByUsernameReq) (*models.GetEncryptedFullInfoRsp, error) {
	encryptedFullInfo, err := s.dbApi.getEncryptedFullInfo(in.Username)
	if err != nil {
		msg := "can not get public info: " + err.Error()
		return &models.GetEncryptedFullInfoRsp{
			Status:            models.GetEncryptedFullInfoRsp_FAIL,
			Message:           &msg,
			EncryptedFullInfo: nil,
		}, nil
	}
	return &models.GetEncryptedFullInfoRsp{
		Status:            models.GetEncryptedFullInfoRsp_OK,
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
