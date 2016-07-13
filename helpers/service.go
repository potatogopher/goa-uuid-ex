package helpers

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	_ "github.com/lib/pq"

	"github.com/potatogopher/goa-uuid-ex/config"
)

// MyNewService sets up the  goa service with default middleware.
func MyNewService() *goa.Service {
	// Create service
	service := goa.New("MyApp")

	service.LogInfo("Starting up", "env", config.App.Environment)

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	return service
}
