package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPeople(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	err := GetPeople(c)

	if err != nil {
		t.Error("Cannot get All people", err)
	}

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetOnePeople(t *testing.T) {

}
