package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	user struct{}
)

func AddUserRoutes(e *echo.Group) {
	ctl := &user{}

	g := e.Group("/users")

	g.POST("", ctl.create)
}

func (ctl *user) create(c echo.Context) error {
	return c.String(http.StatusOK, c.Path())
}
