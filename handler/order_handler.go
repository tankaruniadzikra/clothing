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
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	orderId, err := result.LastInsertId()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	for i := range order.Details {
		if _, err := tx.Exec(createOrderDetail, orderId, order.Details[i].ProductID, order.Details[i].Quantity, order.Details[i].PricePerUnit, order.Details[i].TotalPrice); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}

		if _, err := tx.Exec(decreaseStock, order.Details[i].Quantity, order.Details[i].ProductID); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ReadStocks(db *sql.DB) (map[int]entity.Stock, error) {
	rows, err := db.Query(readStocks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stockMap := make(map[int]entity.Stock, 0)
	for rows.Next() {
		var stock entity.Stock
		err = rows.Scan(&stock.StockID, &stock.ProductID, &stock.Quantity)
		if err != nil {
			return nil, err
		}
		stockMap[stock.ProductID] = stock
	}

	return stockMap, nil
}
