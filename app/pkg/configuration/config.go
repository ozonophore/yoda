package configuration

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	flag "github.com/spf13/pflag"
	"log"
	"os"
)

type Logger struct {
	Level *string
}

type Configuration struct {
	Version   string
	Dsn       string
	BatchSize *int
	SqlLogger Logger
}

var k = koanf.New(".")

func InitConfig() *Configuration {
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	f.StringP("config", "c", "config.yml", "Configuration file")
	f.Parse(os.Args[1:])
	conf, _ := f.GetString("config")
	log.Printf("Initialization config file (%s)", conf)
	k.Load(file.Provider(conf), yaml.Parser())
	var c Configuration
	k.Unmarshal("", &c)
	if c.BatchSize == nil {
		batchSize := 200
		c.BatchSize = &batchSize
	}
	if c.SqlLogger.Level == nil {
		lvl := "Silent"
		c.SqlLogger.Level = &lvl
	}
	return &c
}
