package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPeople(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/person", nil)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)
	c.SetPath("/person")
	err := GetPeople(c)

	if err != nil {
		t.Error("Cannot get All people", err)
	}

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetOnePeople(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	c.SetPath("/person/:id")
	c.SetParamNames("id")
	c.SetParamValues("6213a3d85dc37493c3575a8a")

	err := GetPerson(c)

	if err != nil {
		t.Error("Cannot get this person", err)
	}

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreatePeople(t *testing.T) {
	userJSON := `{"name": "Joe Smith", "height": 195, "weight": 100, "gender": "male"}`

	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	c.SetPath("/person")

	err := CreatePerson(c)
	if err != nil {
		t.Error("Cannot create this person", err)
	}

	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestUpdatePeople(t *testing.T) {
	userJSON := `{"name": "Johnson Smith", "height": 195, "weight": 100, "gender": "male"}`
	e := echo.New()
	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	c.SetPath("/person/:id")
	c.SetParamNames("id")
	c.SetParamValues("6214faf96ad25eff1982b180")

	err := UpdatePerson(c)
	if err != nil {
		t.Error("Cannot update this person", err)
	}

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeletePeople(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodDelete, "/", nil)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	c.SetPath("/person/:id")
	c.SetParamNames("id")
	c.SetParamValues("6214faf96ad25eff1982b180")

	err := DeletePerson(c)
	if err != nil {
		t.Error("Cannot delete this person", err)
	}

	assert.Equal(t, http.StatusNoContent, response.Code)
}
