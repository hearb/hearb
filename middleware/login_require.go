package middleware

import (
	"github.com/labstack/echo"
)

func LoginRequire() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}