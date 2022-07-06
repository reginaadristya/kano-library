package main

import (
	"fmt"
	"kano-library/routes"
	"net/http"
)

type Books struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	RelaseDate string `json:"relaseDate"`
	Length     string `json:"length"`
	Synopsis   string `json:"synopsis"`
}

//connect

func main() {

	//crud.Insert("books", Books{"The Hunger Games : Mockingjay", "Suzzane Collins", "2010", "390 Pages", "asdfghjkl"})
	//crud.Find("books", Books{})
	//crud.Update()
	//crud.Delete()

	http.HandleFunc("/books", routes.Books)
	http.HandleFunc("/books/find", routes.FindByIdBook)
	http.HandleFunc("/books/insert", routes.InsertBook)
	http.HandleFunc("/books/update", routes.UpdateBook)
	http.HandleFunc("/books/delete", routes.DeleteBook)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
