package entity

import "fmt"

// Product represents the Product entity in the database.
type Product struct {
	ProductID   int
	ProductName string
	Description string
	Price       float64
	Material    string
	Weight      float64
	BrandID     int
	BrandName   string
	SizeID      int
	SizeName    string
	ColorID     int
	ColorName   string
	Categories  []Category
}

type ProductCategory struct {
	ProductCategoryID int
	ProductID         int
	CategoryID        int
}

type Size struct {
	SizeID   int
	SizeName string
}

type Sizes []Size

func (s Sizes) ConvertToTable() string {
	var table string
	table = table + ("---\t-------\t\n")
	table = table + ("Id\tUkuran\t\n")
	table = table + ("---\t-------\t\n")

	for i := range s {
		table = table + fmt.Sprintf("%v\t%v\t\n", s[i].SizeID, s[i].SizeName)
	}

	return table
}

type Brand struct {
	BrandID   int
	BrandName string
}

type Brands []Brand

func (b Brands) ConvertToTable() string {
	var table string
	table = table + ("---\t-----\t\n")
	table = table + ("Id\tMerk\t\n")
	table = table + ("---\t-----\t\n")

	for i := range b {
		table = table + fmt.Sprintf("%v\t%v\t\n", b[i].BrandID, b[i].BrandName)
	}

	return table
}

type Color struct {
	ColorID   int
	ColorName string
}

type Colors []Color

func (c Colors) ConvertToTable() string {
	var table string
	table = table + ("---\t-----\t\n")
	table = table + ("Id\tWarna\t\n")
	table = table + ("---\t-----\t\n")

	for i := range c {
		table = table + fmt.Sprintf("%v\t%v\t\n", c[i].ColorID, c[i].ColorName)
	}

	return table
}

type Category struct {
	CategoryID   int
	CategoryName string
}

type Categories []Category

func (c Categories) ConvertToTable() string {
	var table string
	table = table + ("---\t--------\t\n")
	table = table + ("Id\tKategori\t\n")
	table = table + ("---\t--------\t\n")

	for i := range c {
		table = table + fmt.Sprintf("%v\t%v\t\n", c[i].CategoryID, c[i].CategoryName)
	}

	return table
}
