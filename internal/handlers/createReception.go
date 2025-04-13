package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/dataBase"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/service"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
	"github.com/jackc/pgx/v5"
)

func CreateReceptionHandler(w http.ResponseWriter, r *http.Request) {

	userRole := r.Context().Value(models.RoleKey).(string)
	if userRole != "employee" {
		service.ResponseJSON(w, http.StatusForbidden, models.ErrorResponse{
			Errors: "Access denied",
		})

		return
	}

	var req models.ReceptionsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid request body: " + err.Error(),
		})

		return
	}

	receptionID, err := dataBase.DBInProgressReception(req.PVZID)
	if err != nil && err != pgx.ErrNoRows {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid search in progress reception: " + err.Error(),
		})
		return
	}

	if receptionID != "" {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error there is open reception",
		})

		return
	}

	reception, err := dataBase.DBCreateReception(req.PVZID)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid create reception in db: " + err.Error(),
		})

		return
	}

	service.ResponseJSON(w, http.StatusCreated, reception)
}
