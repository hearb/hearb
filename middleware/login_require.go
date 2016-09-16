package middleware

import (
	"net/http"
	"regexp"

	"github.com/hearb/hearb/redis"
	"github.com/labstack/echo"
)

func LoginRequire() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// exclude routes
			if matched, _ := regexp.MatchString(`^\/auth(\/.*)?$`, c.Path()); matched {
				return next(c)
			}

			// check cookie
			cookie, err := c.Cookie("session")
			if err != nil || cookie.Value() == "" {
				// cookie not found
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "login required"})
			}

			// get session from redis
			client := redis.Client()
			userID, err := client.Get(cookie.Value()).Int64()
			if err != nil || userID == 0 {
				// not logged in
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "login required"})
			}

			// store user ID onto context
			c.Set("userID", uint32(userID))

			return next(c)
		}
	}
}
