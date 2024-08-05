package controllers

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports/incoming"
	"github.com/alejandrososa/mars-rover-go/internal/app/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type MoveRoversController struct {
	UseCase incoming.MoveRoversPort
}

func NewMoveRoversController(useCase incoming.MoveRoversPort) *MoveRoversController {
	return &MoveRoversController{UseCase: useCase}
}

func (c *MoveRoversController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Get the "username" variable from the URL
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok || username == "" {
		http.Error(w, "Mission control not found", http.StatusBadRequest)
		return
	}

	var request dto.MoveRoversRequest
	if err := utils.ParseJSONBody(w, r, &request); err != nil {
		// Error handling is already done in ParseJSONBody
		return
	}

	// Set the username in the request DTO
	request.Username = username

	response, err := c.UseCase.Execute(request)
	if err != nil {
		utils.HandleResponse(w, nil, err)
		return
	}
	utils.HandleResponse(w, response, nil)
}
