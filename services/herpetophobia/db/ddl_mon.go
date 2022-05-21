package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"snake/db/connhelper"
)

func createCollection(dbName string, name string, opts ...*options.CreateCollectionOptions) error {
	c := connhelper.GetClient()
	err := c.Database(dbName).CreateCollection(context.TODO(), name, opts...)
	return err
}
