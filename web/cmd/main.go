package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yoda/web/internal/api"
	"github.com/yoda/web/internal/config"
	"github.com/yoda/web/internal/controller"
	_ "github.com/yoda/web/internal/docs"
	"github.com/yoda/web/internal/middleware"
	"github.com/yoda/web/internal/service"
	"github.com/yoda/web/internal/service/auth"
	"github.com/yoda/web/internal/service/dictionary"
	"github.com/yoda/web/internal/service/sale"
	"github.com/yoda/web/internal/service/stock"
	"github.com/yoda/web/internal/storage"
	"os"
)

// @BasePath /rest
// @Host petstore.swagger.io
// @title Swagger Example API
func main() {
	config := config.LoadConfig("config.yml")
	store := storage.NewStorage(config.Database)

	orderService := service.NewOrderService(store)
	saleService := sale.NewSaleService(store)
	authService := auth.NewAuthService(store)
	dictionaryService := dictionary.NewDictionaryService(store)
	stockService := stock.NewStockService(store)

	controller := controller.NewController(store,
		orderService,
		saleService,
		authService,
		dictionaryService,
		stockService,
	)

	e := echo.New()
	e.Use(middleware.JWTValidationMiddleware(authService))
	e.HTTPErrorHandler = middleware.ErrorHandler

	e.Debug = true

	api.RegisterHandlersWithBaseURL(e, controller, "/rest")

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(echoSwagger.URL("/api/openapi.yml")))
	e.Static("/api", "openapi")
	e.Static("/", "public")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}

func initSessionMiddleware(e *echo.Echo) {
	var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	e.Use(session.Middleware(sessionStore))
}
