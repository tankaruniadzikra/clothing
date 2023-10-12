package handler

import (
	"context"
	"database/sql"
	"pair-project/entity"
)

func CreateProduct(db *sql.DB, product entity.Product) error {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	result, err := tx.Exec(insertProduct, product.ProductName, product.Description, product.Price, product.Material, product.Weight, product.BrandID, product.SizeID, product.ColorID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	productId, err := result.LastInsertId()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	for i := range product.Categories {
		_, err := tx.Exec(insertProductCategories, productId, product.Categories[i].CategoryID)
		if err != nil {
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

func ReadProducts(db *sql.DB) ([]entity.Product, error) {
	rows, err := db.Query(readProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []entity.Product{}
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.Description,
			&product.Price,
			&product.Material,
			&product.Weight,
			&product.BrandID,
			&product.BrandName,
			&product.SizeID,
			&product.SizeName,
			&product.ColorID,
			&product.ColorName,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func UpdateProduct(db *sql.DB, product entity.Product) error {
	_, err := db.Exec(updateProduct, product.ProductName, product.Description, product.Price, product.Material, product.Weight, product.BrandID, product.SizeID, product.ColorID, product.ProductID)
	return err
}

func DeleteProduct(db *sql.DB, productID int) error {
	_, err := db.Exec(deleteProduct, productID)
	return err
}

func ReadSizes(db *sql.DB) (entity.Sizes, error) {
	rows, err := db.Query(readSizes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sizes := []entity.Size{}
	for rows.Next() {
		var size entity.Size
		if err := rows.Scan(&size.SizeID, &size.SizeName); err != nil {
			return nil, err
		}
		sizes = append(sizes, size)
	}

	return sizes, nil
}

func ReadBrands(db *sql.DB) (entity.Brands, error) {
	rows, err := db.Query(readBrands)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brands := []entity.Brand{}
	for rows.Next() {
		var brand entity.Brand
		if err := rows.Scan(&brand.BrandID, &brand.BrandName); err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}

	return brands, nil
}

func ReadColors(db *sql.DB) (entity.Colors, error) {
	rows, err := db.Query(readColors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colors := []entity.Color{}
	for rows.Next() {
		var color entity.Color
		if err := rows.Scan(&color.ColorID, &color.ColorName); err != nil {
			return nil, err
		}
		colors = append(colors, color)
	}

	return colors, nil
}

func ReadCategories(db *sql.DB) (entity.Categories, error) {
	rows, err := db.Query(readCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []entity.Category{}
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
