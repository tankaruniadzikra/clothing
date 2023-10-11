-- Insert sample brands for clothing
INSERT INTO Brands (BrandID, BrandName, Description)
VALUES
    (1, 'Adidas', 'Sportswear and athletic shoes'),
    (2, 'Zara', 'Fashion and apparel retailer'),
    (3, 'H&M', 'Global clothing brand');

-- Insert sample sizes for clothing
INSERT INTO Sizes (SizeID, SizeName, Description)
VALUES
    (1, 'Small', 'Small size for clothing'),
    (2, 'Medium', 'Medium size for clothing'),
    (3, 'Large', 'Large size for clothing');

-- Insert sample colors for clothing
INSERT INTO Colors (ColorID, ColorName)
VALUES
    (1, 'Red'),
    (2, 'Blue'),
    (3, 'Black');

-- Insert sample categories for clothing
INSERT INTO Categories (CategoryID, CategoryName, Description)
VALUES
    (1, 'T-Shirts', 'Casual and comfortable'),
    (2, 'Jeans', 'Denim pants'),
    (3, 'Dresses', 'Various dress styles');

-- Insert sample clothing products
INSERT INTO Products (ProductID, ProductName, Description, Price, Material, Weight, BrandID, SizeID, ColorID)
VALUES
    (1, 'Classic T-Shirt', 'A comfortable and classic t-shirt', 20.00, 'Cotton', 0.3, 1, 2, 1),
    (2, 'Slim Fit Jeans', 'Fitted jeans for a modern look', 45.00, 'Denim', 0.7, 2, 1, 2),
    (3, 'Floral Sundress', 'A floral summer dress', 35.00, 'Polyester', 0.4, 3, 3, 1);

-- Associate clothing products with categories
INSERT INTO ProductCategories (ProductCategoryID, ProductID, CategoryID)
VALUES
    (1, 1, 1),
    (2, 2, 2),
    (3, 3, 3);

-- Insert sample stock information for clothing
INSERT INTO Stock (StockID, ProductID, Quantity)
VALUES
    (1, 1, 50),
    (2, 2, 30),
    (3, 3, 20);
