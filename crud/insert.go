package crud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kano-library/models"
	"kano-library/mongo"
	"kano-library/util"
)

//insert data to database

func Insert(collection string, query models.Books) (string, error) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorHandler(err)
	query.Id = primitive.NewObjectID()
	_, err = db.Collection(collection).InsertOne(ctx, query)
	util.ErrorHandler(err)

	return "Insert success", err
}
