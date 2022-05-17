package models

import (
	_ "gorm.io/driver/postgres"
	"kleptophobia/utils"
)

type PrivatePersonRecord struct {
	Username     string `gorm:"primaryKey"`
	PasswordHash []byte
	FirstName    string
}

func ToRecord(person *PrivatePerson) *PrivatePersonRecord {
	return &PrivatePersonRecord{
		Username:     person.Username,
		PasswordHash: utils.GetHash(person.Password),
		FirstName:    person.FirstName,
	}
}
