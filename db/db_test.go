package db

import (
	"fmt"
	"teste-golang/common"
	"teste-golang/types"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOpenDb(t *testing.T) {
	uri, err := common.LoadUri("../local.env")
	if err != nil {
		t.Error("Fail to load uri")
	}

	client, err := OpenDatabase(uri)
	if err != nil {
		t.Error("Fail to open database")
	}
	if client == nil {
		t.Error("Client cannot open")
	}

}

func TestGetCollection(t *testing.T) {
	uri, err := common.LoadUri("../local.env")
	if err != nil {
		t.Error("Fail to load uri")
	}
	collection, err := FindCollection(uri)
	if err != nil {
		t.Error("Fail to get collection", err)
	}

	if collection == nil {
		t.Error("Collection does not exists", err)
	}
}

func TestInsertion(t *testing.T) {
	uri, err := common.LoadUri("../local.env")
	if err != nil {
		t.Error("Fail to load uri")
	}

	var person = types.Person{ID: primitive.NewObjectID(), Name: "Jhon Does", Height: 190, Weight: 85, Gender: "Male"}
	person.IMC = float64(person.Weight) / ((float64(person.Height) / 100) * (float64(person.Height) / 100))

	if err := InsertInCollection(uri, person); err != nil {
		t.Error("Cannot insert data", err)
	}
}

func TestUpdate(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("621015f8c710919c4e49270b")
	if err != nil {
		t.Error("Error, ID doesn't exist", err)
	}
	uri, err := common.LoadUri("../local.env")
	if err != nil {
		t.Error("Fail to load uri")
	}

	var person = types.Person{ID: id, Name: "Fernanda", Height: 155, Weight: 45, Gender: "Female", IMC: (45 / (155 / 10) * (155 / 10))}

	err = UpdatePeople(uri, id, person)
	if err != nil {
		t.Error("Cannot update user", err)
	}

}

func TestDelete(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("6213980380c8ec26213893b7")
	if err != nil {
		t.Error("Error, ID doesn't exist", err)
	}
	uri, err := common.LoadUri("../local.env")
	fmt.Println(uri)
	if err != nil {
		t.Error("Fail to load uri")
	}

	err = DeletePeople(uri, id)
	if err != nil {
		t.Error("Cannot delete user", err)
	}
}
