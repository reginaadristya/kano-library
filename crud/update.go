package crud

import (
	"context"
	"fmt"
	"kano-library/mongo"
	"kano-library/util"
)

// update data from database

type Books struct {
	Title      string `bson:"title"`
	Author     string `bson:"author"`
	RelaseDate string `bson:"relaseDate"`
	Length     string `bson:"length"`
	Synopsis   string `bson:"synopsis"`
}

func UpdateOne(collection string, query map[string]interface{}, update map[string]interface{}) (string, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorHandler(err)

	fmt.Println(update)
	_, err = db.Collection(collection).UpdateOne(ctx, query, update)

	util.ErrorHandler(err)

	return "Update success", err
}
