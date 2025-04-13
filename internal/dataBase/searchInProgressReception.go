package dataBase

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DBInProgressReception(pvzID string) (string, error) {

	query := `
			SELECT id 
			FROM receptions 
			WHERE pvz_id = $1 AND status = 'in_progress'
		`
	var receptionID string

	err := DB.QueryRow(context.Background(), query, pvzID).Scan(&receptionID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", err
		}
		return "", fmt.Errorf("failed to fetch in progress reception ID: %v", err)
	}

	return receptionID, nil
}
