package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/yoda/web/internal/api"
	"net/http"
	"strings"
	"time"
)

const (
	ACCESS_TOKEN = "access_token"
	CLAIM        = "claim"
)

var JwtKey = []byte("my_secret_key")

type Claims struct {
	UserId      int32    `json:"user_id"`
	UserName    string   `json:"user_name"`
	UserEmail   string   `json:"user_email"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

type IAuthService interface {
	CheckUserById(id int32) bool
}

func JWTValidationMiddleware(service IAuthService) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ContextKey: ACCESS_TOKEN,
		SigningKey: JwtKey,
		Skipper: func(c echo.Context) bool {
			return !strings.Contains(c.Path(), "/rest") || c.Request().Header.Get("X-API-KEY") != ""
		},
		ErrorHandler: func(c echo.Context, err error) error {
			message := err.Error()
			return echo.NewHTTPError(http.StatusUnauthorized, api.AuthInfo{
				Success:     false,
				Description: &message,
			})
		},
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			token, err := ParsToken(auth)
			if err != nil {
				return nil, &echojwt.TokenError{Token: token, Err: err}
			}
			claims := token.Claims.(*Claims)
			if claims.ExpiresAt.Time.Before(time.Now()) {
				return nil, &echojwt.TokenError{Token: token, Err: errors.New("token expired")}
			}
			if !service.CheckUserById(claims.UserId) {
				return nil, &echojwt.TokenError{Token: token, Err: errors.New("user not found")}
			}
			c.Set(CLAIM, claims)
			return token, nil
		},
	})
}

func ParsToken(auth string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(auth, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, errors.New("invalid token")
	}
	return token, nil
}
