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

type TelegramBot struct {
	UpdateTimeOut int    `koanf:"update_timeout"`
	LoggingLevel  string `koanf:"logging_level"`
}

type Server struct {
	BaseURL      string `koanf:"base_url"`
	Port         int    `koanf:"port"`
	WriteTimeout int    `koanf:"write_timeout"`
	ReadTimeout  int    `koanf:"read_timeout"`
	IdleTimeout  int    `koanf:"idle_timeout"`
}

type Mq struct {
	Url       string `koanf:"url"`
	Consumer  string `koanf:"consumer"`
	Publisher string `koanf:"publisher"`
	MaxLength int32  `koanf:"max_length"`
}

type Config struct {
	Server       Server      `koanf:"server"`
	LoggingLevel string      `koanf:"logging_level"`
	TelegramBot  TelegramBot `koanf:"telegramBot"`
	Database     Database    `koanf:"database"`
	Mq           Mq          `koanf:"mq"`
}

func fileExists(pathe string) bool {
	_, err := os.Stat(pathe)
	return !os.IsNotExist(err)
}

func LoadConfig(path string) (*Config, error) {
	if err := k.Load(confmap.Provider(map[string]interface{}{
		"server.base_url":            "/api",
		"server.port":                8080,
		"server.write_timeout":       15,
		"server.read_timeout":        15,
		"server.idle_timeout":        60,
		"logging_level":              "INFO",
		"telegramBot.update_timeout": 60,
		"telegramBot.logging_level":  "INFO",
		"database.logging_level":     "Silent",
		"mq.consumer":                "yoda-web-consumer",
		"mq.publisher":               "yoda-server-consumer",
		"mq.max_length":              10,
	}, "."), nil); err != nil {
		return nil, err
	}

	if fileExists(path) {
		if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
			return nil, err
		}
	}
	var c Config
	if err := k.UnmarshalWithConf("", &c, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
		return nil, err
	}
	if err := getEnvs(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

const prefix = "YODA_"

func getEnvs(c *Config) error {
	var err error
	if val, exists := os.LookupEnv(prefix + "SERVER_PORT"); exists {
		c.Server.Port, err = strconv.Atoi(val)
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "SERVER_WRITE_TIMEOUT"); exists {
		c.Server.WriteTimeout, _ = strconv.Atoi(val)
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "SERVER_READ_TIMEOUT"); exists {
		c.Server.ReadTimeout, _ = strconv.Atoi(val)
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "SERVER_IDLE_TIMEOUT"); exists {
		c.Server.IdleTimeout, _ = strconv.Atoi(val)
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "LOGGING_LEVEL"); exists {
		c.LoggingLevel = val
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "TB_UPDATE_TIMEOUT"); exists {
		c.TelegramBot.UpdateTimeOut, _ = strconv.Atoi(val)
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "TB_LOGGING_LEVEL"); exists {
		c.TelegramBot.LoggingLevel = val
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "DB_DSN"); exists {
		c.Database.Dsn = val
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "MQ_URL"); exists {
		c.Mq.Url = val
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "MQ_CONSUMER"); exists {
		c.Mq.Consumer = val
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "MQ_PUBLISHER"); exists {
		c.Mq.Publisher = val
		if err != nil {
			return err
		}
	}
	if val, exists := os.LookupEnv(prefix + "MQ_MAX_LENGTH"); exists {
		v, _ := strconv.Atoi(val)
		c.Mq.MaxLength = int32(v)
		if err != nil {
			return err
		}
	}
	return nil
}
