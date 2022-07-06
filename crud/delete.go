package crud

import (
	"context"
	"kano-library/mongo"
	"kano-library/util"
)

// delete data from database

func DeleteOne(collection string, query map[string]interface{}) (string, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorHandler(err)

	_, err = db.Collection(collection).DeleteOne(ctx, query)

	util.ErrorHandler(err)

	return "Delete success", err
}
