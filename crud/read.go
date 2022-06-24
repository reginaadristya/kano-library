package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kano-library/mongo"
	"kano-library/util"
)

type books struct {
	Name, Author string
}

// read data from database

func Find(collection string, data interface{}) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	csr, err := db.Collection(collection).Find(ctx, bson.D{{}})
	util.ErrorChecker(err)
	defer csr.Close(ctx)

	result := make([]books, 0)
	for csr.Next(ctx) {
		var row books
		err := csr.Decode(&row)
		util.ErrorChecker(err)

		result = append(result, row)
	}

	var i = 0
	for {
		fmt.Println("Name  :", result[i].Name)
		fmt.Println("Author :", result[i].Author)

		i++
	}

	//if len(result) > books {
	//fmt.Println("Name  :", result[0].Name)
	//fmt.Println("Author :", result[0].Author)

	//fmt.Println("Name  :", result[1].Name)
	//fmt.Println("Author :", result[1].Author)

	//fmt.Println("Name  :", result[2].Name)
	//fmt.Println("Author :", result[2].Author)

	//fmt.Println("Name  :", result[3].Name)
	//fmt.Println("Author :", result[3].Author)

	//}

}
