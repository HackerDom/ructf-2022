package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kleptophobia/crypto"
	"kleptophobia/models"
	"kleptophobia/utils"
	"time"
)

type CliClient struct {
	GrpcClient *models.KleptophobiaClient
}

func (cliClient *CliClient) init(config *models.ClientConfig) utils.Closable {
	grpcAddr := fmt.Sprintf("%s:%d", config.GrpcHost, config.GrpcPort)
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	utils.FailOnError(err)

	grpcClient := models.NewKleptophobiaClient(conn)
	cliClient.GrpcClient = &grpcClient

	return conn
}

func buildRegisterRequest() *models.RegisterRequest {
	firstName := utils.ReadValue("First name: ")
	var middleName *string = nil
	if val := utils.ReadValue("Middle name (optional): "); val != "" {
		middleName = &val
	}
	secondName := utils.ReadValue("Second name: ")
	username := utils.ReadValue("Username: ")
	room := utils.ReadIntValue("Room: ")
	diagnosis := utils.ReadValue("Diagnosis: ")

	privatePerson := models.PrivatePerson{
		FirstName:  firstName,
		MiddleName: middleName,
		SecondName: secondName,
		Username:   username,
		Room:       int32(room),
		Diagnosis:  diagnosis,
	}

	password := utils.ReadHiddenValue("Password: ")

	return &models.RegisterRequest{
		Person:   &privatePerson,
		Password: password,
	}
}

func buildGetByUsernameRequest() *models.GetByUsernameRequest {
	return &models.GetByUsernameRequest{
		Username: utils.ReadValue("Username: "),
	}
}

type WithContextType func(ctx context.Context) error

func withDefaultContext(fun WithContextType) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return fun(ctx)
}

func (cliClient *CliClient) Register() error {
	registerRequest := buildRegisterRequest()

	return withDefaultContext(func(ctx context.Context) error {
		registerReply, err := (*cliClient.GrpcClient).Register(ctx, registerRequest)

		if err != nil {
			return errors.New("can not register new user: " + err.Error())
		}

		if registerReply.Status != models.RegisterReply_OK {
			return errors.New("can not register new user: " + registerReply.GetMessage())
		}
		fmt.Println("Success!")
		return nil

	})
}

func (cliClient *CliClient) GetPublicInfo() error {
	getPublicInfoRequest := buildGetByUsernameRequest()

	return withDefaultContext(func(ctx context.Context) error {
		getPublicInfoReply, err := (*cliClient.GrpcClient).GetPublicInfo(ctx, getPublicInfoRequest)

		if err != nil {
			return errors.New("can not get public info: " + err.Error())
		}

		if getPublicInfoReply.Status != models.GetPublicInfoReply_OK {
			return errors.New("can not get public info: " + getPublicInfoReply.GetMessage())
		}

		fmt.Println("\nPublic info: ")
		fmt.Println(proto.MarshalTextString(getPublicInfoReply.GetPerson()))

		return nil
	})
}

func (cliClient *CliClient) GetFullInfo() error {
	getByUsernameRequest := buildGetByUsernameRequest()
	password := utils.ReadHiddenValue("Password: ")

	return withDefaultContext(func(ctx context.Context) error {
		getEncryptedFullInfo, err := (*cliClient.GrpcClient).GetEncryptedFullInfo(ctx, getByUsernameRequest)

		if err != nil {
			return errors.New("can not get full info: " + err.Error())
		}

		if getEncryptedFullInfo.Status != models.GetEncryptedFullInfoReply_OK {
			return errors.New("can not get full info: " + getEncryptedFullInfo.GetMessage())
		}

		encryptedFullInfo := getEncryptedFullInfo.GetEncryptedFullInfo()
		fullInfo := crypto.Decrypt(encryptedFullInfo, utils.GetHash(password))
		var privatePerson models.PrivatePerson

		if err := proto.Unmarshal(fullInfo, &privatePerson); err != nil {
			return errors.New("invalid username or password")
		}

		fmt.Println("\nFull info: ")
		fmt.Println(proto.MarshalTextString(&privatePerson))

		return nil
	})
}

func (cliClient *CliClient) Ping() error {
	return withDefaultContext(func(ctx context.Context) error {
		message := utils.RandString(10)
		pingResponse, err := (*cliClient.GrpcClient).Ping(ctx, &models.PingBody{Message: message})
		if err != nil {
			return err
		}

		if pingResponse == nil {
			return errors.New("ping response is nil")
		}

		if pingResponse.Message != message {
			return errors.New("ping messages are different: " + message + " and " + pingResponse.Message)
		}

		return nil
	})
}
