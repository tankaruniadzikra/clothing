package handler

import (
	"context"
	"database/sql"
	"pair-project/entity"
)

func CreateOrder(db *sql.DB, order entity.Order) error {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	result, err := tx.Exec(createOrder, order.UserID, order.OrderDate, order.TotalAmount)
	if err != nil {
		tx.Rollback()
		return err
	}

	orderId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for i := range order.Details {
		if _, err := tx.Exec(createOrderDetail, orderId, order.Details[i].ProductID, order.Details[i].Quantity, order.Details[i].PricePerUnit, order.Details[i].TotalPrice); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
