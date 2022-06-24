package main

import "kano-library/crud"

type books struct {
	Name   string `bson:"name"`
	Author string `bson:"author"`
}

//connect

func main() {
	//crud.Insert("books", books{Name: "Adventure of Us", Author: "Lang Leav"})
	//crud.Find("books", books{})
	//crud.Update()
	crud.Delete()
}
