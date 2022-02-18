package types

type People struct {
	ID     string `json:"_id" bson:"_id"`
	Name   string
	Heigth int
	Weigth int
	IMC    float64
	Gender string
}
