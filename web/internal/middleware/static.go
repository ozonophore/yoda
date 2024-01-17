package middleware

import (
<<<<<<< Updated upstream
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"io/fs"
	"strings"
)

=======
	"errors"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	indexPage = "index.html"
)

>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
			if strings.Contains(c.Path(), url) {
=======
			if strings.Contains(c.Request().URL.Path, url) {
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
			return echo.StaticFileHandler(config.Root, config.Filesystem)(c)
		}
	}
}
=======
			return StaticFileHandler(config.Root, config.Filesystem)(c)
		}
	}
}

func StaticFileHandler(file string, filesystem fs.FS) echo.HandlerFunc {
	return func(c echo.Context) error {
		return fsFile(c, file, filesystem)
	}
}

func fsFile(c echo.Context, file string, filesystem fs.FS) error {
	f, err := filesystem.Open(file)
	if err != nil {
		return echo.ErrNotFound
	}
	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		url := c.Request().URL.Path
		if !isFileInStatic(url) {
			url = indexPage
		}
		logrus.Debug("Request URL: ", url)
		file = filepath.ToSlash(filepath.Join(file, url)) // ToSlash is necessary for Windows. fs.Open and os.Open are different in that aspect.
		f, err = filesystem.Open(file)
		if err != nil {
			return echo.ErrNotFound
		}
		defer f.Close()
		if fi, err = f.Stat(); err != nil {
			return err
		}
	}
	ff, ok := f.(io.ReadSeeker)
	if !ok {
		return errors.New("file does not implement io.ReadSeeker")
	}
	http.ServeContent(c.Response(), c.Request(), fi.Name(), fi.ModTime(), ff)
	return nil
}

func isFileInStatic(path string) bool {
	if len(path) == 0 || path == "/" {
		return false
	}
	if normalizeFolderName(filepath.Base(path)) == normalizeFolderName(path) && filepath.Ext(path) != "" {
		return true
	}
	baseDir := "/static"
	return strings.HasPrefix(normalizeFolderName(path), baseDir)
}

func normalizeFolderName(foldeerName string) string {
	if strings.HasPrefix(foldeerName, "/") {
		return foldeerName
	}
	return "/" + foldeerName
}
>>>>>>> Stashed changes
