package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/service"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
)

func DummyLoginHandler(w http.ResponseWriter, r *http.Request) {

	var req models.DummyLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid request body: " + err.Error(),
		})

		return
	}

	if req.Role != "employee" && req.Role != "moderator" {
		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "invalid role",
		})

		return
	}

	token, err := service.GenerateJWT(req.Role)
	if err != nil {

		service.ResponseJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Errors: "failed to generate JWT: " + err.Error(),
		})

		return
	}

	service.ResponseJSON(w, http.StatusOK, map[string]string{"token": token})
}
