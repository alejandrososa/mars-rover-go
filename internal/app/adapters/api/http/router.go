package http

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/http/controllers"
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
	healthCheckController := controllers.NewHealthCheckController()

	r.HandleFunc("/health", healthCheckController.CheckHealth).Methods("GET")
}
