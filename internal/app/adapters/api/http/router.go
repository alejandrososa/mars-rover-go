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
	createMissionControlController := config.CreateMissionControlController
	getMissionControlController := config.GetMissionControlController
	moveRoversController := config.MoveRoversController

	r.HandleFunc("/api/health", healthCheckController.HandleRequest).Methods("GET")
	r.HandleFunc("/api/mission-control", createMissionControlController.HandleRequest).Methods("POST")
	r.HandleFunc("/api/mission-control/{username}", getMissionControlController.HandleRequest).Methods("GET")
	r.HandleFunc("/api/mission-control/{username}/move-rovers", moveRoversController.HandleRequest).Methods("POST")
}
