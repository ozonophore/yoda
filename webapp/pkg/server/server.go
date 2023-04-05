package server

import (
	"fmt"
	"github.com/yoda/webapp/pkg/config"
	"net/http"
	"time"
)

func NewServer(config config.Server, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(`0.0.0.0:%d`, config.Port),
		WriteTimeout: time.Second * time.Duration(config.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(config.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(config.IdleTimeout),
		Handler:      router, // Pass our instance of gorilla/mux in.
	}
}
