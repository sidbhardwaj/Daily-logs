package logmanager

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/common/log"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations/health"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations/todo"
	// log "github.com/sidbhardwaj/Daily-logs/log"
)

// Configure this is the router for API calls
func Configure(api *operations.DailyLogsAPI, service Service) {
	api.HealthCheckHealthHandler = health.CheckHealthHandlerFunc(func(params health.CheckHealthParams) middleware.Responder {
		log.Debugf("HealthCheckHealthHandler called")
		return health.NewCheckHealthOK().WithPayload("Working...")
	})

	api.CraeteTODOHandler = todo.CraeteTODOHandlerFunc(func(params todo.CraeteTODOParams) middleware.Responder {
		log.Debugf("CraeteTODOHandlerFunc ...")
		service.
	})
}
