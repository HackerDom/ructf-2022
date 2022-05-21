package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"snake/db"
	"snake/objects"
)

func GetMap(id string) objects.Level {
	res := db.Get(db.DbName, db.ColName, bson.M{"_id": id})
	var level objects.Level
	_ = res.Decode(&level)
	return level
}

func SaveMap(level objects.Level) {
	_, _ = db.InsertDoc(db.DbName, db.ColName, level)
}

func IncCounter(id string) {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$inc", bson.D{{"counter", 1}}}}
	_, _ = db.UpdateDoc(db.DbName, db.ColName, filter, update)
}

func ListId(limit int64, offset int64) objects.Ids {
	if limit >= 10 {
		limit = 10
	}
	opts := options.Find().SetProjection(bson.D{{"_id", 1}}).SetLimit(limit).SetSkip(offset)
	results, _ := db.List(db.DbName, db.ColName, bson.D{}, opts)
	var listId []string
	for _, result := range results {
		mRes := result.Map()
		listId = append(listId, mRes["_id"].(string))
	}
	return objects.Ids{Ids: listId}
}
