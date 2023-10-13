package cli

import (
	"database/sql"
	"fmt"
	"os"
	"pair-project/entity"
	"pair-project/handler"
	"strconv"
	"strings"
	"text/tabwriter"
)

// ProductMenu will prompt the user to select CRUD product
func ProductMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.ProductMenu)
	var option int
	if err := inputScanner(&option, "Pilih menu yang ingin anda akses (1/2/3/4/5): "); err != nil {
		return 0
	}

	switch option {
	case 1:
		if _, err := listProductMenu(db); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
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
		product, err := updateProductInput(db)
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		if err := handler.UpdateProduct(db, product); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		fmt.Printf("Product %s telah berhasil diupdate.\n", product.ProductName)
	case 4:
		productId, err := deleteProductInput()
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		if err := handler.DeleteProduct(db, productId); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			fmt.Println(err)
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

// createProductInput will receive input from CLI and turns them to product entity
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

	colors, err := handler.ReadColors(db)
	if err != nil {
		return
	}
	stringColors := colors.ConvertToTable()
	stringColors = stringColors + "\nMasukkan id warna: "
	err = inputScanner(&product.ColorID, stringColors)
	if err != nil {
		return
	}

	categories, err := handler.ReadCategories(db)
	if err != nil {
		return
	}
	stringCategories := categories.ConvertToTable()
	stringCategories = stringCategories + "\nApabila ingin memasukkan lebih dari satu kategori, gunakan koma, e.g. 1,2,3"
	stringCategories = stringCategories + "\nMasukkan id kategori: "
	var inputCategories string
	err = inputScanner(&inputCategories, stringCategories)
	if err != nil {
		return
	}

	categoryIds := strings.Split(inputCategories, ",")
	for i := range categoryIds {
		categoryId, err := strconv.Atoi(categoryIds[i])
		if err != nil {
			continue
		}

		category := entity.Category{CategoryID: categoryId}
		product.Categories = append(product.Categories, category)
	}

	err = inputScanner(&product.Price, "Masukkan harga product: ")
	if err != nil {
		return
	}

	return
}

// updateProductInput will receive input from CLI and turns them to product entity
func updateProductInput(db *sql.DB) (product entity.Product, err error) {
	err = inputScanner(&product.ProductID, "Masukkan id product yang ingin diperbarui: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.ProductName, "Masukkan nama baru product: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Description, "Masukkan deskripsi baru product: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Material, "Masukkan material baru product: ")
	if err != nil {
		return
	}

	err = inputScanner(&product.Weight, "Masukkan berat baru product: ")
	if err != nil {
		return
	}

	brands, err := handler.ReadBrands(db)
	if err != nil {
		return
	}
	stringBrands := brands.ConvertToTable()
	stringBrands = stringBrands + "\nMasukkan id merk baru: "
	err = inputScanner(&product.BrandID, stringBrands)
	if err != nil {
		return
	}

	sizes, err := handler.ReadSizes(db)
	if err != nil {
		return
	}
	stringSizes := sizes.ConvertToTable()
	stringSizes = stringSizes + "\nMasukkan id ukuran baru: "
	err = inputScanner(&product.SizeID, stringSizes)
	if err != nil {
		return
	}

	colors, err := handler.ReadColors(db)
	if err != nil {
		return
	}
	stringColors := colors.ConvertToTable()
	stringColors = stringColors + "\nMasukkan id warna baru: "
	err = inputScanner(&product.ColorID, stringColors)
	if err != nil {
		return
	}

	err = inputScanner(&product.Price, "Masukkan harga baru product: ")
	if err != nil {
		return
	}

	return
}

// deleteProductInput will receive input from CLI and turns them to product id
func deleteProductInput() (productId int, err error) {
	err = inputScanner(&productId, "Masukkan id product: ")
	if err != nil {
		return
	}

	return
}

// listProductMenu will prints all available products in CLI table
func listProductMenu(db *sql.DB) (map[int]entity.Product, error) {
	products, err := handler.ReadProducts(db)
	if err != nil {
		return nil, err
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "---\t | -----\t | ----------\t | -----\t | ------\t | -------\t | ------\t | ------\t | ------\t")
	fmt.Fprintln(w, "Id \t | Name \t | Deskripsi \t | Merk \t | Warna \t | Ukuran \t | Bahan \t | Berat \t | Harga\t")
	fmt.Fprintln(w, "---\t | -----\t | ----------\t | -----\t | ------\t | -------\t | ------\t | ------\t | ------\t")

	res := make(map[int]entity.Product, 0)
	for i := range products {
		res[products[i].ProductID] = products[i]
		row := fmt.Sprintf("%v\t | %v\t | %v\t | %v\t | %v\t | %v\t | %v\t | %v\t | $%v\t",
			products[i].ProductID,
			products[i].ProductName,
			products[i].Description,
			products[i].BrandName,
			products[i].ColorName,
			products[i].SizeName,
			products[i].Material,
			products[i].Weight,
			products[i].Price)
		fmt.Fprintln(w, row)
	}
	w.Flush()

	return res, nil
}
