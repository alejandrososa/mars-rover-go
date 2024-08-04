package controllers

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/utils"
	"net/http"
)

// HealthCheckController handles health check requests.
type HealthCheckController struct{}

// NewHealthCheckController creates a new instance of HealthCheckController.
func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// HandleRequest handles the HTTP request for the health check endpoint.
// @param w http.ResponseWriter - The response writer to send the HTTP response.
// @param r *http.Request - The HTTP request received.
func (c *HealthCheckController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Create a response indicating the server status is OK.
	response := dto.HealthCheckResponse{Status: "OK"}

	// Use the utility function to send the response.
	utils.HandleResponse(w, response, nil)
}
