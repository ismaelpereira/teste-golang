package api

import (
	"fmt"
	"net/http"
	"teste-golang/common"
	"teste-golang/db"
	"teste-golang/types"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartAPI() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", GetPeople)
	e.POST("/people", CreatePerson)
	e.GET("/person/:id", GetPerson)
	e.PUT("/person/:id", UpdatePerson)
	e.DELETE("/person/:id", DeletePerson)

	e.Logger.Fatal(e.Start(":1323"))

}

func GetPeople(c echo.Context) error {
	uri, err := common.LoadUri("../local.env")
	fmt.Println(uri)
	if err != nil {
		return err
	}
	coll, err := db.FindCollection(uri)
	if err != nil {
		return err
	}

	ctx, cancel := common.Context()
	defer cancel()

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	var peoples []types.Person

	if err = cursor.All(ctx, &peoples); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, peoples)
}

func GetPerson(c echo.Context) error {
	idParam := c.Param("id")

	uri, err := common.LoadUri("../local.env")
	fmt.Println(uri)
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return err
	}

	people, err := db.FindOneInCollection(uri, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, people)
}

func CreatePerson(c echo.Context) error {
	return nil
}

func UpdatePerson(c echo.Context) error {
	return nil
}

func DeletePerson(c echo.Context) error {
	return nil
}
