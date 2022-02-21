package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string
	Heigth int
	Weigth float64
	IMC    float64
	Gender string
}
