package handlers

import (
	"net/http"
	"strings"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/dataBase"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/service"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
)

func DeleteLastProductHandler(w http.ResponseWriter, r *http.Request) {

	userRole := r.Context().Value(models.RoleKey).(string)
	if userRole != "employee" {
		service.ResponseJSON(w, http.StatusForbidden, models.ErrorResponse{
			Errors: "Access denied",
		})

		return
	}

	pvzId := strings.TrimPrefix(r.URL.Path, "/pvz/")
	pvzId = strings.TrimSuffix(pvzId, "/delete_last_product")
	if pvzId == "" {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "error pvz id is missing",
		})
		return
	}

	receptionID, err := dataBase.DBInProgressReception(pvzId)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid search in progress reception: " + err.Error(),
		})

		return
	}

	err = dataBase.DBDeleteLastProduct(receptionID)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid delete product in db: " + err.Error(),
		})

		return
	}

	service.ResponseJSON(w, http.StatusOK, nil)
}
