package logging

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/api"
	integration "github.com/yoda/app/internal/integration/api"
	"net/http"
	"time"
)

type LoggedClientWrapper struct {
	http.Client
	logger *logrus.Logger
}

func (c *LoggedClientWrapper) Do(req *http.Request) (*http.Response, error) {
	c.logger.Debugf("Request: %s %s", req.Method, req.URL.String())
	start := time.Now()
	resp, err := c.Client.Do(req)
	elapsed := time.Since(start)
	c.logger.Debugf("Response: %s %s %s", req.Method, req.URL.String(), elapsed)
	return resp, err
}

// WithLoggerFn allows setting up a logging function, which will be
// wrote a log before sending the request and after.
func WithLoggerFn(level string) api.ClientOption {
	logger := logrus.New()
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(lvl)
	logger.Hooks = logrus.StandardLogger().Hooks
	return func(c *api.Client) error {
		c.Client = &LoggedClientWrapper{
			logger: logger,
		}
		return nil
	}
}

func WithLoggerIntegrationFn(level string) integration.ClientOption {
	logger := logrus.New()
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(lvl)
	logger.Hooks = logrus.StandardLogger().Hooks
	return func(c *integration.Client) error {
		c.Client = &LoggedClientWrapper{
			logger: logger,
		}
		return nil
	}
}

func WithStandardLoggerFn() api.ClientOption {
	return func(c *api.Client) error {
		c.Client = &LoggedClientWrapper{
			logger: logrus.StandardLogger(),
		}
		return nil
	}
}
