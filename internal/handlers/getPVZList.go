package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/dataBase"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/service"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
)

func GetPVZListHandler(w http.ResponseWriter, r *http.Request) {

	userRole := r.Context().Value(models.RoleKey).(string)
	if userRole != "employee" && userRole != "moderator" {
		service.ResponseJSON(w, http.StatusForbidden, models.ErrorResponse{
			Errors: "Access denied",
		})

		return
	}

	var req models.PVZListRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid request body: " + err.Error(),
		})

		return
	}

	var startDate *time.Time
	if req.StartDate != "" {
		startDate, err = service.ParseDate(req.StartDate)
		if err != nil {
			service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
				Errors: "invalid startDate format: " + err.Error(),
			})

			return
		}
	}

	var endDate *time.Time
	if req.EndDate != "" {
		endDate, err = service.ParseDate(req.EndDate)
		if err != nil {
			service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
				Errors: "invalid startDate format: " + err.Error(),
			})

			return
		}
	}

	pvzList, err := dataBase.GetPVZList(startDate, endDate, req.Page, req.Limit)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid get pvz list in db: " + err.Error(),
		})

		return
	}

	service.ResponseJSON(w, http.StatusOK, pvzList)
}
