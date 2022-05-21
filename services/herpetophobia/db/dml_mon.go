package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"snake/db/connhelper"
)

func Get(dbName string, collectionName string, f bson.M, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return getCollection(dbName, collectionName).FindOne(context.TODO(), f, opts...)
}

func List(dbName string, collectionName string, f bson.D, opts ...*options.FindOptions) ([]bson.D, error) {
	ctx := context.TODO()
	cur, err := getCollection(dbName, collectionName).Find(ctx, f, opts...)
	var results []bson.D
	err = cur.All(ctx, &results)
	return results, err
}

func UpdateDocs(dbName string, collectionName string,
	f bson.D, u bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return getCollection(dbName, collectionName).UpdateMany(context.TODO(), f, u, opts...)
}

func UpdateDoc(dbName string, collectionName string,
	f bson.D, u bson.D, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return getCollection(dbName, collectionName).UpdateOne(context.TODO(), f, u, opts...)
}

func InsertDoc(dbName string, collectionName string,
	doc any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {

	one, err := getCollection(dbName, collectionName).InsertOne(context.TODO(), doc, opts...)
	return one, err
}

func DeleteDocs(dbName string, collectionName string, filter any, opts ...*options.DeleteOptions) (int64, error) {
	many, err := getCollection(dbName, collectionName).DeleteMany(context.TODO(), filter, opts...)
	return many.DeletedCount, err
}

func DeleteDoc(dbName string, collectionName string, filter any, opts ...*options.DeleteOptions) (bool, error) {
	one, err := getCollection(dbName, collectionName).DeleteOne(context.TODO(), filter, opts...)
	return one.DeletedCount == 1, err
}

func getCollection(dbName string, collectionName string) *mongo.Collection {
	client := connhelper.GetClient()
	return client.Database(dbName).Collection(collectionName)
}
