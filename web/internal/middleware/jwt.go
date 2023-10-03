package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yoda/web/internal/api"
	"net/http"
	"strings"
)

type JWTSessionConfig struct {
	Skipper middleware.Skipper
}

var DefaultJWTSessionConfig = JWTSessionConfig{
	Skipper: func(c echo.Context) bool {
		return strings.Contains(c.Path(), "/rest/auth/login") || !strings.Contains(c.Path(), "/rest")
	},
}

func JWTSessionWithConfig(config JWTSessionConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultJWTConfig.Skipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			cookie, err := c.Cookie(ACCESS_TOKEN)
			if err != nil {
				message := fmt.Sprintf("Unauthorized: %s", err.Error())
				return c.JSON(http.StatusUnauthorized, &api.AuthInfo{
					Success:     false,
					Description: &message,
				})
			}
			access_token := cookie.Value
			claim := Claims{}
			parsedTokenInfo, err := jwt.ParseWithClaims(access_token, &claim, func(token *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})
			if err != nil {
				message := "Please login again"
				return c.JSON(http.StatusUnauthorized, &api.AuthInfo{
					Success:     false,
					Description: &message,
				})
			}

			if !parsedTokenInfo.Valid {
				message := "Invalid token"
				return c.JSON(http.StatusForbidden, &api.AuthInfo{
					Success:     false,
					Description: &message,
				})
			}

			c.Set(CLAIM, claim)

			return next(c)
		}
	}
}
