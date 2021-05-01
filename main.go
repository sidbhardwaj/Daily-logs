package main

import (
	"github.com/go-openapi/loads"
	"github.com/sidbhardwaj/Daily-logs/cmd"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations"
	ini "github.com/sidbhardwaj/Daily-logs/init"
	"github.com/sidbhardwaj/Daily-logs/logmanager"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	port int
)

func main() {
	initd()
	db := ini.GetDB()
	log.Debugf("main func")
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal("Invalid swagger file ", err)
	}
	api := operations.NewDailyLogsAPI(swaggerSpec)
	logRepo := logmanager.NewRepository(db)
	logService := logmanager.New(logRepo)
	logmanager.Configure(api, logService)

	if err := cmd.Start(api, port); err != nil {
		logrus.Fatal("Failed to start", err)
	}
}

func initd() {
	port = 8080
}
