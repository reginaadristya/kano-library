package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kano-library/mongo"
	"kano-library/util"
)

// update data from database

func Update() {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	var selector = bson.M{"name": "Adventure of Us"}
	var changes = books{"Love", "Lang Leav"}
	_, err = db.Collection("books").UpdateOne(ctx, selector, bson.M{"$set": changes})
	util.ErrorChecker(err)

	fmt.Println("Update success!")
}
