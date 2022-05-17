package main

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	pb "kleptophobia/models"
)

type DBApi struct {
	db *gorm.DB
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (dbapi *DBApi) init(host string, port int, dbUsername, password, dbname string) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, dbUsername, password, dbname)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlconn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	failOnError(err)

	failOnError(db.AutoMigrate(&pb.PrivatePersonRecord{}))
	dbapi.db = db
}

func (dbapi *DBApi) register(person *pb.PrivatePerson) error {
	privatePersonRecord := pb.ToRecord(person)
	result := dbapi.db.Create(&privatePersonRecord)

	if result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				return errors.New(fmt.Sprintf("can not register: username %s is already exists", person.Username))
			}
			return errors.New("can not register: " + result.Error.Error())
		}
	}
	return nil
}
