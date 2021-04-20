package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.GET("/corpse/:name")
func (a *Application) getCorpse(c echo.Context) error {
	name := c.Param("name")
	corpse, err := a.model.GetCorpse(name)
	if err != nil {
		a.errorLog.Print(err)
	}
	return c.String(http.StatusOK, fmt.Sprint(corpse))
}
