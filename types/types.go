package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string
	Heigth int
	Weigth int
	IMC    float64
	Gender string
}
