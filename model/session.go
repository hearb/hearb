package model

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/hearb/hearb/redis"
	"github.com/labstack/echo"
)

type (
	// Session is cookie-based session relation
	Session struct {
		SID    string
		UserID uint32
	}
)

const (
	sessionAge = 7 * 24 * time.Hour
)

// RegenerateSessionID regenerates session ID
func (s *Session) RegenerateSessionID() {
	raw := make([]byte, 32)
	rand.Read(raw)
	s.SID = base64.StdEncoding.EncodeToString(raw)
}

// Save stores session data onto redis & cookie
func (s *Session) Save(c echo.Context) error {
	// save to redis
	client := redis.Client()
	if err := client.Set(s.SID, s.UserID, sessionAge).Err(); err != nil {
		return err
	}

	// save to cookie
	cookie := &echo.Cookie{}
	cookie.SetName("session")
	cookie.SetValue(s.SID)
	cookie.SetPath("/")
	cookie.SetExpires(time.Now().Add(sessionAge))
	cookie.SetSecure(true)
	cookie.SetHTTPOnly(true)
	c.SetCookie(cookie)

	return nil
}
