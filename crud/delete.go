package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kano-library/mongo"
	"kano-library/util"
)

// delete data from database

func Delete() {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	var selector = bson.M{"name": "Adventure of Us"}
	_, err = db.Collection("books").DeleteOne(ctx, selector)
	util.ErrorChecker(err)

	fmt.Println("Delete success!")
}
