package dataBase

import (
	"context"
	"fmt"
)

func DBDeleteLastProduct(receptionID string) error {

	tx, err := DB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("error to start transaction: %v", err)
	}
	defer tx.Rollback(context.Background())
	var UUID string
	query := `
			WITH last_product AS (
    			SELECT
        			id,
        			product_id[array_length(product_id, 1)] AS last_product_uuid
    			FROM receptions
    			WHERE id = $1 
			)
			UPDATE receptions
			SET product_id = array_remove(product_id, (SELECT last_product_uuid FROM last_product))
			WHERE id = $1
			RETURNING (SELECT last_product_uuid FROM last_product);
		`

	err = tx.QueryRow(context.Background(), query, receptionID).Scan(&UUID)
	if err != nil || UUID == "" {
		return fmt.Errorf("failed to fetch product: %v", err)
	}

	exec := `
			DELETE FROM products
			WHERE id = $1;
		`
	_, err = tx.Exec(context.Background(), exec, UUID)
	if err != nil {
		return fmt.Errorf("error exec delete product: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("error commit transaction: %v", err)
	}

	return nil
}
