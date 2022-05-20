package models

import (
	"errors"

	"google.golang.org/protobuf/proto"
	_ "gorm.io/driver/postgres"

	"kleptophobia/crypto"
	"kleptophobia/utils"
)

type PersonRecord struct {
	Username               string `gorm:"primaryKey;type:varchar(30)"`
	PasswordHash           []byte
	FirstName              string `gorm:"type:varchar(30)"`
	SecondName             string `gorm:"type:varchar(30)"`
	Room                   uint32
	EncryptedPrivatePerson []byte
}

func PrivatePersonToRecord(person *PrivatePerson, password string) (*PersonRecord, error) {
	data, err := proto.Marshal(person)
	utils.FailOnError(err)

	passwordHash := utils.GetHash(password)
	c := crypto.NewCipher(passwordHash)
	encryptedPrivatePerson, err := c.Encrypt(data)
	if err != nil {
		return nil, errors.New("can not encrypt data: " + err.Error())
	}

	return &PersonRecord{
		Username:               person.Username,
		PasswordHash:           passwordHash,
		FirstName:              person.FirstName,
		SecondName:             person.SecondName,
		Room:                   person.Room,
		EncryptedPrivatePerson: encryptedPrivatePerson,
	}, nil
}

func PersonRecordToPublic(person *PersonRecord) *PublicPerson {
	return &PublicPerson{
		FirstName:  person.FirstName,
		SecondName: person.SecondName,
		Username:   person.Username,
		Room:       person.Room,
	}
}
