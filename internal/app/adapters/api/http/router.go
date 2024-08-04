package http

import (
	"github.com/alejandrososa/mars-rover-go/internal/config"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	r := &Router{
		Router: mux.NewRouter(),
	}
	r.SetupRoutes()
	return r
}

func (r *Router) SetupRoutes() {
	healthCheckController := config.HealthCheckController
	createPlatformController := config.CreatePlatformController

	r.HandleFunc("/health", healthCheckController.HandleRequest).Methods("GET")
	r.HandleFunc("/create-mission-control", createPlatformController.HandleRequest).Methods("POST")
	// Add other routes and controllers here
}
