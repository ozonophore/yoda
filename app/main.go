package main

import (
	"github.com/yoda/app/pkg/client"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/db"
	"log"
)

func main() {
	c := configuration.InitConfig()
	d := db.InitDatabase(c.Dsn)
	prvd := db.NewDbProvider(d)
	var lstn client.EventListener = prvd
	clnt := client.NewWBClient()
	err := clnt.Parsing(&lstn)
	if err != nil {
		log.Panicf("Error after parsing: %s", err)
	}
}
