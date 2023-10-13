package handler

const (
	readProducts = `
SELECT
	P.ProductID,
	P.ProductName,
	P.Description,
	P.Price,
	P.Material,
	P.Weight,
	P.BrandID,
	COALESCE(B.BrandName, ''),
	P.SizeID,
	COALESCE(S.SizeName, ''),
	P.ColorID,
	COALESCE(C.ColorName, '')
FROM
	Products P
LEFT JOIN
	Brands B ON P.BrandID = B.BrandID
LEFT JOIN
	Sizes S ON P.SizeID = S.SizeID
LEFT JOIN
	Colors C ON P.ColorID = C.ColorID
GROUP BY
	P.ProductID`

	insertProduct = `INSERT INTO Products (ProductName, Description, Price, Material, Weight, BrandID, SizeID, ColorID) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	insertProductCategories = `INSERT INTO ProductCategories (ProductID, CategoryID) VALUES (?, ?)`

	updateProduct = `UPDATE Products SET ProductName = ?, Description = ?, Price = ?, Material = ?, Weight = ?, BrandID = ?, SizeID = ?, ColorID = ? WHERE ProductID = ?`

	deleteProduct = `DELETE FROM Products WHERE ProductID = ?`
)

const (
	readSizes = `SELECT SizeID, SizeName FROM Sizes`

	readBrands = `SELECT BrandID, BrandName FROM Brands`

	readColors = `SELECT ColorID, ColorName FROM Colors`

	readCategories = `SELECT CategoryID, CategoryName FROM Categories`
)
