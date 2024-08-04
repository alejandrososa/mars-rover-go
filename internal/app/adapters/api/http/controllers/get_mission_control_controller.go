package controllers

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports/incoming"
	"github.com/alejandrososa/mars-rover-go/internal/app/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// GetMissionControlController handles requests to get mission control details for a user.
type GetMissionControlController struct {
	UseCase incoming.GetMissionControlPort
}

// NewGetMissionControlController creates a new GetMissionControlController.
func NewGetMissionControlController(useCase incoming.GetMissionControlPort) *GetMissionControlController {
	return &GetMissionControlController{UseCase: useCase}
}

// HandleRequest handles the HTTP request for getting mission control details.
func (c *GetMissionControlController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Get the "username" variable from the URL
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok || username == "" {
		http.Error(w, "Mission control not found", http.StatusBadRequest)
		return
	}

	// Create the request DTO
	request := dto.GetMissionControlRequest{Username: username}

	response, err := c.UseCase.Execute(request)
	if err != nil {
		utils.HandleResponse(w, nil, err)
		return
	}
	utils.HandleResponse(w, response, nil)
}
