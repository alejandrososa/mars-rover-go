package controllers

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports/incoming"
	"github.com/alejandrososa/mars-rover-go/internal/app/utils"
	"net/http"
)

// CreateMissionControlController handles requests for creating a new mission control.
type CreateMissionControlController struct {
	// UseCase is the use case for creating mission control, following the port pattern.
	UseCase incoming.CreateMissionControlPort
}

// NewCreateMissionControlController initializes a new controller with the provided use case.
func NewCreateMissionControlController(useCase incoming.CreateMissionControlPort) *CreateMissionControlController {
	return &CreateMissionControlController{UseCase: useCase}
}

// HandleRequest processes the HTTP request for creating a mission control.
// @param w http.ResponseWriter - The response writer to write the HTTP response.
// @param r *http.Request - The HTTP request received.
func (c *CreateMissionControlController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateMissionControlRequest

	// Parse the JSON body of the request into the request DTO.
	if err := utils.ParseJSONBody(w, r, &request); err != nil {
		// Handle any errors that occurred during parsing.
		return
	}

	// Execute the use case with the parsed request data.
	response, err := c.UseCase.Execute(request)
	if err != nil {
		// Handle the error response.
		utils.HandleResponse(w, nil, err)
		return
	}

	// Handle the successful response.
	utils.HandleResponse(w, response, nil)
}
