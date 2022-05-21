package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"snake/db/connhelper"
	"time"
)

var DbName = os.Getenv("dbName")
var ColName = os.Getenv("collectionName")

func Migrate() {
	if connhelper.ConnUrl == "" || DbName == "" || ColName == "" {
		log.Fatal("Service need specified 'ME_CONFIG_MONGODB_URL', 'dbName' and 'collectionName' in .env")
	}
	err := createCollection(DbName, ColName)
	if err != nil {
		switch err.(type) {
		case mongo.CommandError:
			log.Printf("Database '%s' and collection '%s' already exists", DbName, ColName)
		default:
			log.Fatal(err)
		}
	}
	createIndexes(DbName, ColName)
}

func createIndexes(dbName string, collectionName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	c := getCollection(dbName, collectionName)
	_, err := c.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"createdAt", 1}},
		Options: options.Index().SetExpireAfterSeconds(15.5 * 60)})
	specs, _ := c.Indexes().ListSpecifications(ctx)
	if len(specs) > 2 {
		log.Printf("Too many indexes in the collection '%s'", ColName)
	}
	if err != nil {
		var commandErr mongo.CommandError
		if !errors.As(err, &commandErr) {
			log.Fatal(err)
		}
	}
	cancel()
}
