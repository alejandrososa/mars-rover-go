package config

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/http/controllers"
)

var HealthCheckController *controllers.HealthCheckController

func init() {
	HealthCheckController = controllers.NewHealthCheckController()
}
