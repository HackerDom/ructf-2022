package objects

import "time"

type Level struct {
	Id        string    `bson:"_id"`
	Secret    string    `bson:"secret"`
	Counter   int       `bson:"counter"`
	Init      [256]byte `bson:"init"`
	Flag      string    `bson:"flag"`
	CreatedAt time.Time `bson:"createdAt"`
}
