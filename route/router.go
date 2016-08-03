package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

var app *echo.Echo

// Init server handling
func Init() {
	app = echo.New()
}

// Run server
func Run() {
	app.Run(standard.New(":5556"))
}
