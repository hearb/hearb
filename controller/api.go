package controller

import (
	"github.com/hearb/hearb/api/v1"
	mw "github.com/hearb/hearb/middleware"
	"github.com/labstack/echo"
)

func AddAPIRoutes(e *echo.Echo) {
	g := e.Group("/api/v1", mw.LoginRequire())
	v1.AddUserRoutes(g)
}
