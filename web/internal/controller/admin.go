package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/middleware"
	"net/http"
)

func (c *Controller) Profile(ctx echo.Context) error {
	claims := ctx.Get(middleware.CLAIM).(*middleware.Claims)
	permissions := make([]api.Permission, len(claims.Permissions))
	for i, permission := range claims.Permissions {
		permissions[i] = api.Permission(permission)
	}
	profile := &api.Profile{
		Email:       claims.UserEmail,
		Name:        claims.UserName,
		Permissions: permissions,
	}
	return ctx.JSON(http.StatusOK, profile)
}

func (c *Controller) Refresh(ctx echo.Context) error {
	return nil
}

func (c *Controller) Login(ctx echo.Context) error {
	var login api.LoginInfo
	err := ctx.Bind(&login)
	if err != nil {
		description := fmt.Sprintf("Unauthorized %s", err.Error())
		return ctx.JSON(http.StatusUnauthorized, api.AuthInfo{
			Success:     false,
			Description: &description,
		})
	}
	tokenString, expirationTime, err := c.authService.CreateToken(&login)
	if err != nil {
		description := fmt.Sprintf("Unauthorized %s", err.Error())
		return ctx.JSON(http.StatusUnauthorized, api.AuthInfo{
			Success:     false,
			Description: &description,
		})
	}
	cookie := &http.Cookie{
		Name:    middleware.ACCESS_TOKEN,
		Value:   tokenString,
		Expires: expirationTime,
	}
	ctx.SetCookie(cookie)
	return ctx.JSON(http.StatusOK, api.AuthInfo{
		Success:     true,
		AccessToken: &tokenString,
	})
}
