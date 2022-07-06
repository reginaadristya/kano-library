package crud

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kano-library/models"
	"kano-library/mongo"
	"kano-library/util"
)

type books struct {
	Title, Author, RelaseDate, Synopsis string
}

// read data from database

func Find(collection string, data interface{}) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorHandler(err)

	csr, err := db.Collection(collection).Find(ctx, bson.D{{}})
	util.ErrorHandler(err)
	defer csr.Close(ctx)

	result := make([]books, 0)
	for csr.Next(ctx) {
		var row books
		err := csr.Decode(&row)
		util.ErrorHandler(err)

		result = append(result, row)
	}

	var i = 0
	for {
		fmt.Println("KANO LIBRARY BOOKS")
		fmt.Println("Title  :", result[i].Title)
		fmt.Println("Author :", result[i].Author)
		fmt.Println("Relase Date  :", result[i].RelaseDate)
		fmt.Println("Sypnosis :", result[i].Author)
		fmt.Println(" ")
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
func FindAll(collection string) ([]byte, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorHandler(err)

	csr, err := db.Collection(collection).Find(ctx, bson.M{})
	util.ErrorHandler(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		util.ErrorHandler(err)

		result = append(result, row)

	}
	data, err := json.Marshal(result)
	util.ErrorHandler(err)
	return data, err
}

//read data by specific ID

func FindOne(collection string, query map[string]interface{}) ([]byte, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorHandler(err)

	csr, err := db.Collection(collection).Find(ctx, query)
	util.ErrorHandler(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		util.ErrorHandler(err)

		result = append(result, row)

	}
	data, err := json.Marshal(result)
	util.ErrorHandler(err)
	return data, err
}
