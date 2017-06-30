package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// RootGet just a empty main url
func RootGet(c echo.Context) error {
	mainPage := `Hey, this is a main page.
Auth: LukeEuler
Email: luke16times@gmail.com
Have a good day!`
	return c.String(http.StatusOK, mainPage)
}
