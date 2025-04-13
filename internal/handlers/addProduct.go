package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/dataBase"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/service"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
)

func AddProductHandler(w http.ResponseWriter, r *http.Request) {

	userRole := r.Context().Value(models.RoleKey).(string)
	if userRole != "employee" {
		service.ResponseJSON(w, http.StatusForbidden, models.ErrorResponse{
			Errors: "Access denied",
		})

		return
	}

	var req models.ProductsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid request body: " + err.Error(),
		})

		return
	}

	receptionID, err := dataBase.DBInProgressReception(req.PVZID)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid search in progress reception: " + err.Error(),
		})

		return
	}

	product, err := dataBase.DBAddProduct(req.Type, receptionID)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid create add product in db: " + err.Error(),
		})

		return
	}

	service.ResponseJSON(w, http.StatusCreated, product)
}
