package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/dataBase"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/service"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
)

func CreatePVZHandler(w http.ResponseWriter, r *http.Request) {

	userRole := r.Context().Value(models.RoleKey).(string)
	if userRole != "moderator" {
		service.ResponseJSON(w, http.StatusForbidden, models.ErrorResponse{
			Errors: "Access denied",
		})

		return
	}

	var req models.PVZRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid request body: " + err.Error(),
		})

		return
	}

	if !service.IsValidCity(req.City) {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid city for a pvz establishment",
		})

		return
	}

	pvz, err := dataBase.DBCreatePVZ(req.City)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid create pvz in db: " + err.Error(),
		})

		return
	}

	service.ResponseJSON(w, http.StatusCreated, pvz)
}
