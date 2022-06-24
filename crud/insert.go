package crud

import (
	"context"
	"fmt"
	"kano-library/mongo"
	"kano-library/util"
)

//insert data to database

func Insert(collection string, data interface{}) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	_, err = db.Collection(collection).InsertOne(ctx, data)
	util.ErrorChecker(err)
	fmt.Println("Inserted Successfully")
}
