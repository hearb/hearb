package controller

import (
	"github.com/hearb/hearb/api"
	mw "github.com/hearb/hearb/middleware"
	"github.com/labstack/echo"
)

func AddAPIRoutes(e *echo.Echo) {
	g := e.Group("/api/v1", mw.LoginRequire())
	api.AddUserRoutes(g)
}
