package configuration

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"os"
)

type Logger struct {
	Level string
}

type Database struct {
	Dsn          string `koanf:"dsn"`
	LoggingLevel string `koanf:"logging_level"`
}

type Config struct {
	Version      string   `koanf:"version"`
	Database     Database `koanf:"database"`
	LoggingLevel string   `koanf:"logging_level"` //panic | fatal | error | warn | info | debug | trace
	Token        string   `koanf:"token"`
}

var conf = koanf.Conf{
	Delim:       ".",
	StrictMerge: true,
}

var k = koanf.NewWithConf(conf)

func fileExists(pathe string) bool {
	_, err := os.Stat(pathe)
	return !os.IsNotExist(err)
}

func InitConfig(path string) *Config {
	k.Load(confmap.Provider(map[string]interface{}{
		"version":                "0.0.1",
		"logging_level":          "info",
		"database.logging_level": "silent",
	}, "."), nil)
	if fileExists(path) {
		k.Load(file.Provider(path), yaml.Parser())
	}
	var c Config
	k.UnmarshalWithConf("", &c, koanf.UnmarshalConf{Tag: "koanf"})
	getEnvs(&c)
	return &c
}

func getEnvs(c *Config) {
	if v, exists := os.LookupEnv("YODA_VERSION"); exists {
		c.Version = v
	}
	if v, exists := os.LookupEnv("YODA_DATABASE_DSN"); exists {
		c.Database.Dsn = v
	}
	if v, exists := os.LookupEnv("YODA_DATABASE_LOGGING_LEVEL"); exists {
		c.Database.LoggingLevel = v
	}
	if v, exists := os.LookupEnv("YODA_LOGGING_LEVEL"); exists {
		c.LoggingLevel = v
	}
	if v, exists := os.LookupEnv("YODA_TOKEN"); exists {
		c.Token = v
	}
}
