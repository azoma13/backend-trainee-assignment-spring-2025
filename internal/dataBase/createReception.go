package dataBase

import (
	"context"
	"fmt"
	"time"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
	"github.com/google/uuid"
)

func DBCreateReception(pvzID string) (models.Reception, error) {

	dateTime := time.Now()
	UUID := uuid.New()

	exec := `
			INSERT INTO receptions 
				(id, date_time, pvz_id, status) 
				VALUES ($1, $2, $3, $4)
		`

	_, err := DB.Exec(context.Background(), exec, UUID, dateTime, pvzID, "in_progress")
	if err != nil {
		return models.Reception{}, fmt.Errorf("error exec insert new reception: %v", err)
	}

	return models.Reception{
		ID:       UUID.String(),
		DateTime: dateTime.String(),
		PVZID:    pvzID,
		Status:   "in_progress",
	}, nil
}
