package dataBase

import (
	"context"
	"fmt"
)

func DBSearchPVZID(pvzID string) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM pvz 
			WHERE id = $1
		)
	`
	var found bool

	err := DB.QueryRow(context.Background(), query, pvzID).Scan(&found)
	if err != nil {
		return false, fmt.Errorf("failed to check pvz ID existence: %v", err)
	}

	if !found {
		return false, fmt.Errorf("pvzID not found")
	}

	return found, nil
}
