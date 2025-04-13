package dataBase

import (
	"context"
	"fmt"
	"time"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
	"github.com/google/uuid"
)

func DBCreatePVZ(city string) (models.PVZ, error) {

	registrationDate := time.Now()
	UUID := uuid.New()

	exec := `
			INSERT INTO pvz 
				(id, registration_date, city) 
				VALUES ($1, $2, $3)
		`
	_, err := DB.Exec(context.Background(), exec, UUID, registrationDate, city)
	if err != nil {
		return models.PVZ{}, fmt.Errorf("error exec insert new pvz: %v", err)
	}

	return models.PVZ{
		ID:               UUID.String(),
		RegistrationDate: registrationDate.String(),
		City:             city,
	}, nil
}
