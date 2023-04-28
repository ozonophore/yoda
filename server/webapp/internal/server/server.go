package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yoda/webapp/internal/api"
	"github.com/yoda/webapp/internal/config"
	"github.com/yoda/webapp/internal/controller"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func StartServer(config config.Server, logger *logrus.Logger) *Server {
	cntrl := controller.NewServerApi(logger)

	router := api.HandlerWithOptions(cntrl, api.GorillaServerOptions{
		BaseURL: "/api",
	})

	srv := Server{server: &http.Server{
		Addr:         fmt.Sprintf(`0.0.0.0:%d`, config.Port),
		WriteTimeout: time.Second * time.Duration(config.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(config.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(config.IdleTimeout),
		Handler:      router, // Pass our instance of gorilla/mux in.
	}}

	srv.server.RegisterOnShutdown(func() {
		logger.Info("Http Server shutdown")
	})
	go func() {
		if err := srv.server.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()
	logger.Infof("Server started on port %d", config.Port)
	return &srv
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
