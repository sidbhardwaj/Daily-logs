// +build !aws_lambda

package cmd

import (
	"github.com/sidbhardwaj/Daily-logs/gen/restapi"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations"
)

func Start(api *operations.DailyLogsAPI, port int) error {
	server := restapi.NewServer(api)
	defer server.Shutdown() // nolint
	server.Port = port
	return server.Serve()
}
