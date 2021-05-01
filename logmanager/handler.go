package logmanager

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations/health"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations/todo"
	log "github.com/sirupsen/logrus"
	// log "github.com/sidbhardwaj/Daily-logs/log"
)

// Configure this is the router for API calls
func Configure(api *operations.DailyLogsAPI, service Service) {
	api.HealthCheckHealthHandler = health.CheckHealthHandlerFunc(func(params health.CheckHealthParams) middleware.Responder {
		log.Debug("HealthCheckHealthHandler called")
		return health.NewCheckHealthOK().WithPayload("Working...")
	})

	api.TodoCraeteTODOHandler = todo.CraeteTODOHandlerFunc(func(params todo.CraeteTODOParams) middleware.Responder {
		log.Debug("CraeteTODOHandlerFunc ...")
		result, err := service.CreateTODO(params.HTTPRequest.Context(), params)
		if err != nil {
			log.Error(err)
			return todo.NewCraeteTODOBadRequest().WithPayload(&todo.CraeteTODOBadRequestBody{
				Code:    "400",
				Message: fmt.Sprintf("ERR: %+v", err),
			})
		}
		return todo.NewCraeteTODOOK().WithPayload(result)
	})
	api.TodoListTODOHandler = todo.ListTODOHandlerFunc(func(params todo.ListTODOParams) middleware.Responder {
		log.Debug("ListTODOHandlerFunc ...")
		result, err := service.ListTODO(params.HTTPRequest.Context(), params)
		if err != nil {
			log.Error(err)
			return todo.NewListTODOBadRequest().WithPayload(&todo.ListTODOBadRequestBody{
				Code:    "400",
				Message: fmt.Sprintf("ERR: %+v", err),
			})
		}
		return todo.NewListTODOOK().WithPayload(result)
	})
}
