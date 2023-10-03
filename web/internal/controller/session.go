package controller

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const SESSION_NAME = "session"

func GetFromSession(ctx echo.Context, key string) interface{} {
	sess, _ := session.Get(SESSION_NAME, ctx)
	if sess == nil {
		return nil
	}
	return sess.Values[key]
}

type HttpSession struct {
	sess *sessions.Session
	ctx  echo.Context
}

func GetSession(ctx echo.Context) *HttpSession {
	sess, _ := session.Get(SESSION_NAME, ctx)
	return &HttpSession{
		sess: sess,
		ctx:  ctx,
	}
}

func (s *HttpSession) SetToSession(key string, value interface{}) *HttpSession {
	s.sess.Values[key] = value
	return s
}

func (s *HttpSession) Save() {
	s.sess.Save(s.ctx.Request(), s.ctx.Response())
}

func RemoveSession(ctx echo.Context) {
	sess, _ := session.Get(SESSION_NAME, ctx)
	sess.Options.MaxAge = -1
	sess.Save(ctx.Request(), ctx.Response())
}
