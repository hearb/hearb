package route

import (
	"time"

	"github.com/hearb/hearb/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type (
	application struct {
		engine *echo.Echo
	}
)

func calcAgeInSecond(years, months, days int) int {
	now := time.Now()
	dur := now.AddDate(years, months, days).Sub(now)
	return int(dur.Seconds())
}

// Init server handling
func Init() *application {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength:  24,
		TokenLookup:  "header:X-CSRF-Token",
		ContextKey:   "csrf",
		CookieName:   "csrf",
		CookiePath:   "/",
		CookieMaxAge: 5 * 60, // 5分
		CookieSecure: true,
	}))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            calcAgeInSecond(0, 6, 0), // 6ヶ月
		ContentSecurityPolicy: "default-src: 'self'",
	}))

	controller.AddAuthRoutes(e)
	controller.AddAPIRoutes(e)

	return &application{
		engine: e,
	}
}

// Run server
func (app *application) Run() {
	app.engine.Run(standard.New(":5556"))
}
