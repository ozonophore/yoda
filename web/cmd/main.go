package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/config"
	stock "github.com/yoda/web/internal/controller"
	_ "github.com/yoda/web/internal/docs"
	"github.com/yoda/web/internal/storage"
	"strings"
)

func main() {
	config := config.LoadConfig("config.yml")
	store := storage.NewStorage(config.Database)
	srv := stock.NewServer(store)

	e := echo.New()
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/swagger") || strings.Contains(c.Path(), "/api")
		},
		KeyLookup: "header:X-API-KEY",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == "9609c3c0-2026-11ee-be56-0242ac120002", nil
		},
	}))
	e.Debug = true

	api.RegisterHandlers(e, srv)

	e.Static("/api", "openapi")
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(echoSwagger.URL("/api/openapi.yml")))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
