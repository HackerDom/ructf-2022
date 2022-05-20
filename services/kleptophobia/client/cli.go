package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"kleptophobia/config"
	"kleptophobia/crypto"
	"kleptophobia/models"
	"kleptophobia/utils"
)

type CliClient struct {
	GrpcClient *models.KleptophobiaClient
}

func (cliClient *CliClient) init(config *config.ClientConfig) utils.Closable {
	grpcAddr := fmt.Sprintf("%s:%d", config.GrpcHost, config.GrpcPort)
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	utils.FailOnError(err)

	grpcClient := models.NewKleptophobiaClient(conn)
	cliClient.GrpcClient = &grpcClient

	return conn
}

func buildRegisterReq() *models.RegisterReq {
	firstName := utils.ReadValueWithValidation("First name: ", models.NameRegex)
	middleName := utils.ReadValueWithValidation("Middle name: ", models.NameRegex)
	secondName := utils.ReadValueWithValidation("Second name: ", models.NameRegex)
	username := utils.ReadValueWithValidation("Username: ", models.UsernameRegex)
	room := utils.ReadUIntValue("Room: ")
	diagnosis := utils.ReadValueWithValidation("Diagnosis: ", models.DiagnosisRegex)

	privatePerson := models.PrivatePerson{
		FirstName:  firstName,
		MiddleName: middleName,
		SecondName: secondName,
		Username:   username,
		Room:       room,
		Diagnosis:  diagnosis,
	}

	password := utils.ReadHiddenValue("Password: ")

	return &models.RegisterReq{
		Person:   &privatePerson,
		Password: password,
	}
}

func buildGetByUsernameReq() *models.GetByUsernameReq {
	return &models.GetByUsernameReq{
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
	registerReq := buildRegisterReq()

	return withDefaultContext(func(ctx context.Context) error {
		registerRsp, err := (*cliClient.GrpcClient).Register(ctx, registerReq)

		if err != nil {
			return errors.New("can not register new user: " + err.Error())
		}

		if registerRsp.Status != models.RegisterRsp_OK {
			return errors.New("can not register new user: " + registerRsp.GetMessage())
		}
		fmt.Println("Success!")
		return nil

	})
}

func (cliClient *CliClient) GetPublicInfo() error {
	getPublicInfoReq := buildGetByUsernameReq()

	return withDefaultContext(func(ctx context.Context) error {
		getPublicInfoRsp, err := (*cliClient.GrpcClient).GetPublicInfo(ctx, getPublicInfoReq)

		if err != nil {
			return errors.New("can not get public info: " + err.Error())
		}

		if getPublicInfoRsp.Status != models.GetPublicInfoRsp_OK {
			return errors.New("can not get public info: " + getPublicInfoRsp.GetMessage())
		}

		fmt.Println("\nPublic info: ")
		fmt.Println(proto.MarshalTextString(getPublicInfoRsp.GetPerson()))

		return nil
	})
}

func (cliClient *CliClient) GetFullInfo() error {
	getByUsernameReq := buildGetByUsernameReq()
	password := utils.ReadHiddenValue("Password: ")

	return withDefaultContext(func(ctx context.Context) error {
		getEncryptedFullInfo, err := (*cliClient.GrpcClient).GetEncryptedFullInfo(ctx, getByUsernameReq)

		if err != nil {
			return errors.New("can not get full info: " + err.Error())
		}

		if getEncryptedFullInfo.Status != models.GetEncryptedFullInfoRsp_OK {
			return errors.New("can not get full info: " + getEncryptedFullInfo.GetMessage())
		}

		c := crypto.NewCipher(utils.GetHash(password))

		encryptedFullInfo := getEncryptedFullInfo.GetEncryptedFullInfo()
		fullInfo, err := c.Decrypt(encryptedFullInfo)
		if err != nil {
			return err
		}
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
