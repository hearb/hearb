package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	auth struct{}
)

func AddAuthRoutes(e *echo.Echo) {
	ctl := &auth{}

	g := e.Group("/auth")

	g.GET("", ctl.get)
	g.GET("/callback", ctl.callback)
}

func (ctl *auth) get(c echo.Context) error {
	return c.String(http.StatusOK, c.Path())
}

func (ctl *auth) callback(c echo.Context) error {
	return c.String(http.StatusOK, c.Path())
}
