package dataBase

import (
	"context"
	"fmt"
	"time"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
)

func GetPVZList(startDate, endDate *time.Time, page, limit int) ([]models.PVZListResponse, error) {

	offset := (page - 1) * limit

	query := `
			SELECT id, registration_date::text, city
			FROM pvz
			WHERE ($1::TIMESTAMP IS NULL OR registration_date >= $1)
	  			AND ($2::TIMESTAMP IS NULL OR registration_date <= $2)
			ORDER BY registration_date DESC
			LIMIT $3 OFFSET $4
		`
	rows, err := DB.Query(context.Background(), query, startDate, endDate, limit, offset)
	if err != nil {
		return []models.PVZListResponse{}, fmt.Errorf("failed to fetch PVZ list: %v", err)
	}
	defer rows.Close()

	var res []models.PVZListResponse
	for rows.Next() {

		var pvz models.PVZ
		if err := rows.Scan(&pvz.ID, &pvz.RegistrationDate, &pvz.City); err != nil {
			return []models.PVZListResponse{}, fmt.Errorf("failed to scan PVZ row: %v", err)
		}

		var getPVZ models.PVZListResponse
		getPVZ.PVZ = pvz

		getPVZ.Reception, err = getReceptionsForPVZ(pvz.ID)
		if err != nil {
			return []models.PVZListResponse{}, fmt.Errorf("failed to fetch receptions for PVZ: %v", err)
		}
		res = append(res, getPVZ)
	}

	if res == nil {
		return []models.PVZListResponse{}, fmt.Errorf("not found elements on the page")
	}

	return res, nil
}

func getReceptionsForPVZ(pvzId string) ([]models.GetReception, error) {

	query := `
        	SELECT id, date_time::text, status
        	FROM receptions
        	WHERE pvz_id = $1
        	ORDER BY date_time DESC
    	`

	rows, err := DB.Query(context.Background(), query, pvzId)
	if err != nil {
		return nil, fmt.Errorf("failed query reception: %v", err)
	}
	defer rows.Close()

	var receptions []models.GetReception
	receptions = []models.GetReception{}
	for rows.Next() {

		var reception models.GetReception
		reception.PVZID = pvzId

		err = rows.Scan(&reception.ID, &reception.DateTime, &reception.Status)
		if err != nil {
			return []models.GetReception{}, fmt.Errorf("failed to scan reception row: %v", err)
		}

		reception.Products, err = getProductForReception(reception.ID)
		if err != nil {
			return []models.GetReception{}, fmt.Errorf("failed get product: %v", err)
		}
		receptions = append(receptions, reception)
	}

	return receptions, nil
}

func getProductForReception(receptionId string) ([]models.GetProduct, error) {

	query := `
    		SELECT id, date_time::text, type
        	FROM products
        	WHERE reception_id = $1
        	ORDER BY date_time ASC
    	`

	rows, err := DB.Query(context.Background(), query, receptionId)
	if err != nil {
		return []models.GetProduct{}, fmt.Errorf("failed query product: %v", err)
	}
	defer rows.Close()

	var products []models.GetProduct
	products = []models.GetProduct{}
	for rows.Next() {

		var product models.GetProduct

		if err := rows.Scan(&product.ID, &product.DateTime, &product.Type); err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}

		products = append(products, product)
	}

	return products, nil
}
