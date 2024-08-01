package controllers

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/utils"
	"net/http"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (c *HealthCheckController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	response := dto.HealthCheckResponse{Status: "OK"}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}
