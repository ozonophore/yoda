package configuration

import (
	"errors"
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"os"
	"strconv"
)

type Logger struct {
	Level string
}

// Configuration for the ozon
type Ozon struct {
	Host string `koanf:"host"`
}

// Configuration for the wb
type Wb struct {
	Host          string `koanf:"host"`
	RemainingDays int    `koanf:"remaining_days"` // days to load stocks
}

// Configuration for the orderd
type Order struct {
	LoadedDays int `koanf:"loaded_days"`
}

type Database struct {
	Dsn          string `koanf:"dsn"`
	LoggingLevel string `koanf:"logging_level"`
}

type Mq struct {
	Url       string `koanf:"url"`
	Consumer  string `koanf:"consumer"`
	Publisher string `koanf:"publisher"`
	MaxLength int32  `koanf:"max_length"`
}

type Integration struct {
	Timeout  int    `koanf:"timeout"`
	Token    string `koanf:"token"`
	Host     string `koanf:"host"`
	LogLevel string `koanf:"logging_level"` //panic | fatal | error | warn | info | debug | trace
}

type Config struct {
	Version      string      `koanf:"version"`
	Database     Database    `koanf:"database"`
	BatchSize    int         `koanf:"batch_size"`
	Timeout      int         `koanf:"timeout"`
	Mq           Mq          `koanf:"mq"`
	Order        Order       `koanf:"order"`
	Wb           Wb          `koanf:"wb"`
	Ozon         Ozon        `koanf:"ozon"`
	LoggingLevel string      `koanf:"logging_level"` //panic | fatal | error | warn | info | debug | trace
	Integration  Integration `koanf:"integration"`
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
		"version":                   "0.0.1",
		"batch_size":                500,
		"timeout":                   60,
		"logging_level":             "info",
		"database.logging_level":    "silent",
		"mq.consumer":               "yoda-server-consumer",
		"mq.publisher":              "yoda-publisher",
		"mq.max_length":             10,
		"order.loaded_days":         30,
		"wb.host":                   "https://statistics-api.wildberries.ru",
		"wb.remaining_days":         30,
		"ozon.host":                 "https://api-seller.ozon.ru",
		"integration.timeout":       60,
		"integration.token":         "ZTc4NzJkODYtYjQ4Mi00MTA5LWI3OTMtNTI4NWNiMjI1OGEw",
		"integration.host":          "http://d4440c0ccab0.sn.mynetname.net:8899/ERP_IIS/hs/Market",
		"integration.logging_level": "info",
	}, "."), nil)
	if fileExists(path) {
		k.Load(file.Provider(path), yaml.Parser())
	}
	var c Config
	k.UnmarshalWithConf("", &c, koanf.UnmarshalConf{Tag: "koanf"})
	if err := getEnvs(&c); err != nil {
		errors.New(fmt.Sprintf("Error while getting envs: %s", err))
	}
	return &c
}

func getEnvs(c *Config) error {
	var err error
	if v, exists := os.LookupEnv("YODA_VERSION"); exists {
		c.Version = v
	}
	if v, exists := os.LookupEnv("YODA_DATABASE_DSN"); exists {
		c.Database.Dsn = v
	}
	if v, exists := os.LookupEnv("YODA_DATABASE_LOGGING_LEVEL"); exists {
		c.Database.LoggingLevel = v
	}
	if v, exists := os.LookupEnv("YODA_MQ_URL"); exists {
		c.Mq.Url = v
	}
	if v, exists := os.LookupEnv("YODA_MQ_CONSUMER"); exists {
		c.Mq.Consumer = v
	}
	if v, exists := os.LookupEnv("YODA_MQ_PUBLISHER"); exists {
		c.Mq.Publisher = v
	}
	if v, exists := os.LookupEnv("YODA_BATCH_SIZE"); exists {
		if c.BatchSize, err = strconv.Atoi(v); err != nil {
			return err
		}
	}
	if v, exists := os.LookupEnv("YODA_TIMEOUT"); exists {
		if c.Timeout, err = strconv.Atoi(v); err != nil {
			return err
		}
	}
	if v, exists := os.LookupEnv("YODA_ORDER_LOADED_DAYS"); exists {
		if c.Order.LoadedDays, err = strconv.Atoi(v); err != nil {
			return err
		}
	}
	if v, exists := os.LookupEnv("YODA_WB_HOST"); exists {
		c.Wb.Host = v
	}
	if v, exists := os.LookupEnv("YODA_WB_REMAINING_DAYS"); exists {
		if c.Wb.RemainingDays, err = strconv.Atoi(v); err != nil {
			return err
		}
	}
	if v, exists := os.LookupEnv("YODA_OZON_HOST"); exists {
		c.Ozon.Host = v
	}
	if v, exists := os.LookupEnv("YODA_LOGGING_LEVEL"); exists {
		c.LoggingLevel = v
	}
	if v, exists := os.LookupEnv("YODA_INTEGRATION_TIMEOUT"); exists {
		if c.Integration.Timeout, err = strconv.Atoi(v); err != nil {
			return err
		}
	}
	if v, exists := os.LookupEnv("YODA_INTEGRATION_TOKEN"); exists {
		c.Integration.Token = v
	}
	if v, exists := os.LookupEnv("YODA_INTEGRATION_HOST"); exists {
		c.Integration.Host = v
	}
	if v, exists := os.LookupEnv("YODA_INTEGRATION_LOGGING_LEVEL"); exists {
		c.Integration.LogLevel = v
	}
	return nil
}
