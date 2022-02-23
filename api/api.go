package api

import (
	"fmt"
	"net/http"
	"teste-golang/common"
	"teste-golang/db"
	"teste-golang/rabbit"
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

	e.GET("/person", getPeople)
	e.POST("/person", createPerson)
	e.GET("/person/:id", getPerson)
	e.PUT("/person/:id", updatePerson)
	e.DELETE("/person/:id", deletePerson)

	e.Logger.Fatal(e.Start(":1323"))

}

func getPeople(c echo.Context) error {
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

func getPerson(c echo.Context) error {
	idParam := c.Param("id")

	uri, err := common.LoadUri("../local.env")
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

func createPerson(c echo.Context) error {
	p := new(types.Person)

	if err := c.Bind(p); err != nil {
		return err
	}
	p.ID = primitive.NewObjectID()
	p.IMC = float64(p.Weight) / ((float64(p.Height) / 100) * (float64(p.Height) / 100))

	uri, err := common.LoadUri("../local.env")
	if err != nil {
		return err
	}
	if err := db.InsertInCollection(uri, *p); err != nil {
		return err
	}

	message := "Cadastro de pessoa" + p.ID.String() + ":" + p.Name
	if err := rabbit.SendMessage(message); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, p)
}

func updatePerson(c echo.Context) error {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)

	p := new(types.Person)
	if err := c.Bind(p); err != nil {
		return err
	}
	p.ID = id
	p.IMC = float64(p.Weight) / ((float64(p.Height) / 100) * (float64(p.Height) / 100))

	if err != nil {
		return err
	}
	uri, err := common.LoadUri("../local.env")
	if err != nil {
		return err
	}
	if err := db.UpdatePeople(uri, id, *p); err != nil {
		return err
	}

	message := "Edição da pessoa" + p.ID.String() + ":" + p.Name
	if err := rabbit.SendMessage(message); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func deletePerson(c echo.Context) error {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return err
	}
	uri, err := common.LoadUri("../local.env")
	if err != nil {
		return err
	}
	err = db.DeletePeople(uri, id)
	if err != nil {
		return err
	}

	message := "Exclusão da pessoa" + idParam
	if err := rabbit.SendMessage(message); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
