package routes

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kano-library/crud"
	"kano-library/models"
	"kano-library/util"
	"net/http"
)

//get all books data

func Books(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := crud.FindAll("books")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// get spesific books with id

func FindByIdBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := crud.FindOne("books", bson.M{"_id": objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// insert books
func InsertBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := crud.Insert("books", book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// update books
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PATCH" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		util.ErrorHandler(err)
		result, err := crud.UpdateOne("books", bson.M{"_id": objId}, bson.M{"$set": bson.M{"title": book.Title, "author": book.Author, "relaseDate": book.RelaseDate, "length": book.Length, "synopsis": book.Synopsis}})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		requestId := r.URL.Query().Get("_id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := crud.DeleteOne("books", bson.M{"id": objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
