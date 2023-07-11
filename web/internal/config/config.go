package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"os"
	"strconv"
)

var conf = koanf.Conf{
	Delim:       ".",
	StrictMerge: true,
}
var k = koanf.NewWithConf(conf)

type Database struct {
	Dsn          string `koanf:"dsn"`
	LoggingLevel string `koanf:"logging_level"`
}

type Config struct {
	Database Database `koanf:"database"`
	Port     int      `koanf:"port"`
}

func fileExists(pathe string) bool {
	_, err := os.Stat(pathe)
	return !os.IsNotExist(err)
}

func LoadConfig(path string) *Config {
	if err := k.Load(confmap.Provider(map[string]interface{}{
		"database.logging_level": "Silent",
		"port":                   8080,
	}, "."), nil); err != nil {
		panic(err)
	}

	if fileExists(path) {
		if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
			panic(err)
		}
	}
	var c Config
	if err := k.UnmarshalWithConf("", &c, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
		panic(err)
	}
	if err := getEnvs(&c); err != nil {
		panic(err)
	}
	return &c
}

const prefix = "YODA_"

func getEnvs(c *Config) error {
	var err error
	if val, exists := os.LookupEnv(prefix + "DB_DSN"); exists {
		c.Database.Dsn = val
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "PORT"); exists {
		c.Port, err = strconv.Atoi(val)
		if err != nil {
			return err
		}
	}
	return nil
}
