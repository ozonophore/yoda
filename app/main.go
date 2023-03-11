package main

import (
	"fmt"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/app/pkg/service"
	"log"
)

func main() {
	config := configuration.InitConfig()
	database := repository.InitDatabase(config)

	jobId := 1
	repository := repository.NewRepositoryDAO(database)
	wbService := service.NewWBService("OWNER", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NJRCI6IjFiMzVmODljLTMyNGYtNGM3OS05NzhhLTkwMmYwODk3Mjc4YiJ9.WeYv1vqA46_9D5up2LRUeSBZCXxSBNcmH8lUhG9Jii0", config)
	err := wbService.Parsing(repository, &jobId)
	if err != nil {
		log.Panicf("Error after parsing: %s", err)
	}

	ozonService := service.NewOzonService("OWNER", "538358", "8539be7e-a37f-4b4f-b5e1-3879e5f1738c", config)

	err = ozonService.Parsing(repository, &jobId)
	if err != nil {
		log.Panicf("Error after parsing: %s", err)
	}
	fmt.Println("#-----------#")
}
