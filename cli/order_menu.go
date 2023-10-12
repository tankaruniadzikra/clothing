package cli

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
	"pair-project/handler"
	"time"
)

func OrderMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.OrderMenu)

	order, err := createOrderInput(db)
	if err != nil {
		fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
		fmt.Println("Kembali ke menu...")
		return 0
	}

	if err := handler.CreateOrder(db, order); err != nil {
		fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
		fmt.Println("Kembali ke menu...")
		return 0
	}

	fmt.Println("Order telah berhasil disimpan.")

	return 0
}

func createOrderInput(db *sql.DB) (order entity.Order, err error) {
	order.TotalAmount, order.Details, err = selectProducts(db)
	if err != nil {
		return
	}

	fmt.Println("Berikut adalah user yang sudah terdaftar di Cakra Clothing Store.")
	err = listUserMenu(db)
	if err != nil {
		return
	}

	err = inputScanner(&order.UserID, "Masukkan id customer untuk order ini: ")
	if err != nil {
		return
	}

	order.OrderDate = time.Now().Format("2006-01-02")

	return
}

func selectProducts(db *sql.DB) (float64, []entity.OrderDetail, error) {
	var totalAmount float64
	var orderDetails []entity.OrderDetail
	fmt.Println("Berikut adalah product yang dapat digunakan untuk proses order.")

	productMap, err := listProductMenu(db)
	if err != nil {
		fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
		fmt.Println("Kembali ke menu...")
		return totalAmount, orderDetails, err
	}

	stockMap, err := handler.ReadStocks(db)
	if err != nil {
		return totalAmount, orderDetails, err
	}

	var choice string
	for {
		orderDetail := entity.OrderDetail{}
		err = inputScanner(&orderDetail.ProductID, "Masukkan product id yang ingin ditambahkan: ")
		if err != nil {
			return totalAmount, orderDetails, err
		}

		err = inputScanner(&orderDetail.Quantity, "Masukkan jumlah product yang ingin ditambahkkan: ")
		if err != nil {
			return totalAmount, orderDetails, err
		}

		if stockMap[orderDetail.ProductID].Quantity < orderDetail.Quantity {
			fmt.Println("Product terpilih tidak dapat ditambahkan.")
			fmt.Println("Silahkan pilih product lain.")
			continue
		}

		// soft decrease stock, similar to add to cart
		productStock := stockMap[orderDetail.ProductID]
		productStock.Quantity = productStock.Quantity - orderDetail.Quantity
		stockMap[orderDetail.ProductID] = productStock

		// adding more info to orderDetail
		orderDetail.PricePerUnit = productMap[orderDetail.ProductID].Price
		orderDetail.TotalPrice = orderDetail.PricePerUnit * float64(orderDetail.Quantity)
		totalAmount = totalAmount + orderDetail.TotalPrice
		orderDetails = append(orderDetails, orderDetail)

		fmt.Println("Product telah berhasil ditambahkan!")

		err = inputScanner(&choice, "Tambahkan product lain? (y/n): ")
		if err != nil {
			return totalAmount, orderDetails, err
		}

		if choice == "y" || choice == "yes" {
			continue
		} else {
			break
		}
	}

	return totalAmount, orderDetails, nil
}
