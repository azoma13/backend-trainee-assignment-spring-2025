package dataBase

import (
	"context"
	"fmt"
	"strings"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
)

func DBCloseLastReception(receptionID string) (models.Reception, error) {
	query := `
		UPDATE receptions
		SET status = 'close'
		WHERE id = $1
		RETURNING 
			id::text,
            to_char(date_time, 'YYYY-MM-DD HH24:MI:SS') AS date_time,
            pvz_id::text,
            status::text,
			array_to_string(product_id, ',') AS product_id_string;
	`
	var reception models.Reception
	var productsID string

	err := DB.QueryRow(context.Background(), query, receptionID).Scan(&reception.ID, &reception.DateTime, &reception.PVZID, &reception.Status, &productsID)
	if err != nil {
		return models.Reception{}, fmt.Errorf("failed to close reception: %v", err)
	}

	if productsID != "" {
		reception.Products = strings.Split(productsID, ",")
	} else {
		reception.Products = []string{}
	}

	return reception, nil
}
