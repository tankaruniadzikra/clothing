package cli

import (
	"database/sql"
	"fmt"
	"os"
	"pair-project/entity"
	"pair-project/handler"
)

func ProductMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.ProductMenu)
	var option int
	if err := inputScanner(&option, "Pilih menu yang ingin anda akses (1/2/3/4/5): "); err != nil {
		return 0
	}

	switch option {
	case 1:
		products, err := handler.ReadProducts(db)
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		for i := range products {
			// TODO: find nice ways to print product list
			fmt.Println(products[i])
		}
	case 2:
		product, err := createProductInput(db)
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		if err := handler.CreateProduct(db, product); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		fmt.Printf("Product %s telah berhasil disimpan.\n", product.ProductName)
	case 3:
		product := updateProductInput(db)
		if err := handler.UpdateProduct(db, product); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		fmt.Printf("Product %s telah berhasil diupdate.\n", product.ProductName)
	case 4:
		productId := deleteProductInput()
		if err := handler.DeleteProduct(db, productId); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		fmt.Println("Product telah berhasil dihapus.")
	case 5:
		fmt.Println("Terima kasih telah mengunjungi Cakra Clothing Store. Sampai jumpa!")
		os.Exit(0)
		return 2
	}

	return 0
}

func createProductInput(db *sql.DB) (product entity.Product, err error) {
	err = inputScanner(&product.ProductName, "Masukkan nama product: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Description, "Masukkan deskripsi product: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Material, "Masukkan material product: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Weight, "Masukkan berat product: ")
	if err != nil {
		return
	}

	brands, err := handler.ReadBrands(db)
	if err != nil {
		return
	}

	stringBrands := brands.ConvertToTable()
	stringBrands = stringBrands + "\nMasukkan id merk: "
	err = inputScanner(&product.BrandID, stringBrands)
	if err != nil {
		return
	}

	sizes, err := handler.ReadSizes(db)
	if err != nil {
		return
	}

	stringSizes := sizes.ConvertToTable()
	stringSizes = stringSizes + "\nMasukkan id ukuran: "
	err = inputScanner(&product.SizeID, stringSizes)
	if err != nil {
		return
	}

	err = inputScanner(&product.ColorID, "Masukkan id warna: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Categories, "Masukkan id kategori. Apabila lebih dari satu, pisahkan dengan koma: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Price, "Masukkan harga product: ")
	if err != nil {
		return
	}

	return
}

func updateProductInput(db *sql.DB) entity.Product {
	return entity.Product{}
}

func deleteProductInput() int {
	return 0
}
