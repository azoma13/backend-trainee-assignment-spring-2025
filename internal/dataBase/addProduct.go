package dataBase

import (
	"context"
	"fmt"
	"time"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
	"github.com/google/uuid"
)

func DBAddProduct(productType, receptionID string) (models.Product, error) {

	tx, err := DB.Begin(context.Background())
	if err != nil {
		return models.Product{}, fmt.Errorf("error to start transaction: %v", err)
	}
	defer tx.Rollback(context.Background())

	dateTime := time.Now()
	UUID := uuid.New()

	exec := `
			INSERT INTO products 
				(id, date_time, type, reception_id) 
				VALUES ($1, $2, $3, $4)
		`

	_, err = tx.Exec(context.Background(), exec, UUID, dateTime, productType, receptionID)
	if err != nil {
		return models.Product{}, fmt.Errorf("error exec insert product: %v", err)
	}

	exec = `
			UPDATE receptions 
				SET product_id = array_append(product_id, $1)
				WHERE id = $2;
		`
	_, err = tx.Exec(context.Background(), exec, UUID, receptionID)
	if err != nil {
		return models.Product{}, fmt.Errorf("error exec update receptions: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return models.Product{}, fmt.Errorf("error commit transaction: %v", err)
	}

	return models.Product{
		ID:          UUID.String(),
		DateTime:    dateTime.String(),
		Type:        productType,
		ReceptionID: receptionID,
	}, nil
}
