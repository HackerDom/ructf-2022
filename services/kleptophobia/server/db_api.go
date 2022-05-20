package main

import (
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"kleptophobia/config"
	"kleptophobia/models"
	"kleptophobia/utils"
)

type DBApi struct {
	db *gorm.DB
}

func (dbapi *DBApi) init(pgConfig *config.PGConfig) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pgConfig.Host, pgConfig.Port, pgConfig.Username, pgConfig.Password, pgConfig.DbName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlconn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	utils.FailOnError(err)

	utils.FailOnError(db.AutoMigrate(models.PersonRecord{}))
	dbapi.db = db
}

func (dbapi *DBApi) register(person *models.PrivatePerson, password string) error {
	if err := models.ValidatePrivatePerson(person); err != nil {
		return err
	}
	privatePersonRecord, err := models.PrivatePersonToRecord(person, password)
	if err != nil {
		return errors.New("can not translate private person to record: " + err.Error())
	}
	result := dbapi.db.Create(&privatePersonRecord)

	if result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				return errors.New(fmt.Sprintf("username %s is already exists", person.Username))
			}
			return result.Error
		}
	}
	return nil
}

func (dbapi *DBApi) getPublicInfo(username string) (*models.PersonRecord, error) {
	var person models.PersonRecord

	if res := dbapi.db.Take(&person, "username = ?", username); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("can not found user with username = " + username)
	}

	return &person, nil
}

func (dbapi *DBApi) getEncryptedFullInfo(username string) ([]byte, error) {
	person, err := dbapi.getPublicInfo(username)
	if err != nil {
		return nil, err
	}
	return person.EncryptedPrivatePerson, nil
}
