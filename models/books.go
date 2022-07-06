package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Books struct {
	Id         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title"`
	Author     string             `json:"author"`
	RelaseDate string             `json:"relaseDate"`
	Length     string             `json:"length"`
	Synopsis   string             `json:"synopsis"`
}
