package models

import (
	"kleptophobia/crypto"
	"kleptophobia/utils"

	"google.golang.org/protobuf/proto"
	_ "gorm.io/driver/postgres"
)

type PersonRecord struct {
	Username               string `gorm:"primaryKey"`
	PasswordHash           []byte
	FirstName              string
	SecondName             string
	Room                   int32
	EncryptedPrivatePerson []byte
}

func PrivatePersonToRecord(person *PrivatePerson, password string) *PersonRecord {
	data, err := proto.Marshal(person)
	utils.FailOnError(err)

	passwordHash := utils.GetHash(password)
	c := crypto.NewCipher(passwordHash)
	encryptedPrivatePerson, _ := c.Encrypt(data)

	return &PersonRecord{
		Username:               person.Username,
		PasswordHash:           passwordHash,
		FirstName:              person.FirstName,
		SecondName:             person.SecondName,
		Room:                   person.Room,
		EncryptedPrivatePerson: encryptedPrivatePerson,
	}
}

func PersonRecordToPublic(person *PersonRecord) *PublicPerson {
	return &PublicPerson{
		FirstName:  person.FirstName,
		SecondName: person.SecondName,
		Username:   person.Username,
		Room:       person.Room,
	}
}
