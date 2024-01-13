package middleware

import (
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"io/fs"
	"strings"
)

type (
	// StaticConfig defines the config for Static middleware.
	StaticConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware2.Skipper

		// Root directory from where the static content is served.
		// Required.
		Root string `yaml:"root"`

		// Index file for serving a directory.
		// Optional. Default value "index.html".
		Index string `yaml:"index"`

		// Enable HTML5 mode by forwarding all not-found requests to root so that
		// SPA (single-page application) can handle the routing.
		// Optional. Default value false.
		HTML5 bool `yaml:"html5"`

		// Enable directory browsing.
		// Optional. Default value false.
		Browse bool `yaml:"browse"`

		// Enable ignoring of the base of the URL path.
		// Example: when assigning a static middleware to a non root path group,
		// the filesystem path is not doubled
		// Optional. Default value false.
		IgnoreBase bool `yaml:"ignoreBase"`

		// Filesystem provides access to the static content.
		// Optional. Defaults to http.Dir(config.Root)
		Filesystem fs.FS `yaml:"-"`
	}
)

var (
	// DefaultStaticConfig is the default Static middleware config.
	DefaultStaticConfig = StaticConfig{
		Root:  "./public/",
		Index: "index.html",
	}
)

func Static(filesystem fs.FS, skipedUrls ...string) echo.MiddlewareFunc {
	c := DefaultStaticConfig
	c.Filesystem = filesystem
	c.Skipper = func(c echo.Context) bool {
		for _, url := range skipedUrls {
			if strings.Contains(c.Path(), url) {
				return true
			}
		}
		return false
	}
	return StaticWithConfig(c)
}

func StaticWithConfig(config StaticConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Root == "" {
		config.Root = "." // For security we want to restrict to CWD.
	}
	if config.Skipper == nil {
		config.Skipper = middleware2.DefaultStaticConfig.Skipper
	}
	if config.Index == "" {
		config.Index = middleware2.DefaultStaticConfig.Index
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}
			return echo.StaticFileHandler(config.Root, config.Filesystem)(c)
		}
	}
}
