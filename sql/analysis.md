# Tables

1. **Users Table:**
Columns:
    - UserID (Primary Key, INT)
    - Email (VARCHAR(100), NOT NULL)
    - Password (VARCHAR(255), NOT NULL)
    - FirstName (VARCHAR(50))
    - LastName (VARCHAR(50))
Relationships:
    - This table is self-contained and does not have direct relationships with other tables.
Analysis:
    - The Users table is used to manage user accounts and store user-related information.
    - UserID serves as the primary key.
    - Email and Password are used for user authentication.
    - FirstName and LastName are optional for storing additional user details.

2. **Orders Table:**
Columns:
    - OrderID (Primary Key, INT)
    - UserID (Foreign Key, INT)
    - OrderDate (DATE, NOT NULL)
    - TotalAmount (DECIMAL(10, 2), NOT NULL)
Relationships:
    - Many-to-One relationship with Users (UserID references Users.UserID).
Analysis:
    - The Orders table is used to store information about orders made by users.
    - OrderID is the primary key.
    - UserID establishes a relationship with the Users table to associate orders with specific users.
    - OrderDate records the date and time the order was placed.
    - TotalAmount represents the total cost of the order.

3. **OrderItems Table:**
Columns:
    - OrderItemID (Primary Key, INT)
    - OrderID (Foreign Key, INT)
    - ProductID (Foreign Key, INT)
    - Quantity (INT, NOT NULL)
    - PricePerUnit (DECIMAL(10, 2), NOT NULL)
    - TotalPrice (DECIMAL(10, 2), NOT NULL)
Relationships:
    - Many-to-One relationship with Orders (OrderID references Orders.OrderID).
    - Many-to-One relationship with Products (ProductID references Products.ProductID).
Analysis:
    - The OrderItems table tracks individual items within orders.
    - OrderItemID is the primary key.
    - OrderID associates items with specific orders.
    - ProductID links items to products available in the Products table.
    - Quantity, PricePerUnit, and TotalPrice record item-specific details.

4. **Products Table:**
Columns:
    - ProductID (Primary Key, INT)
    - ProductName (VARCHAR(255), NOT NULL)
    - Description (TEXT)
    - Price (DECIMAL(10, 2), NOT NULL)
    - Material (VARCHAR(255))
    - Weight (DECIMAL(8, 2))
    - BrandID (Foreign Key, INT)
    - SizeID (Foreign Key, INT)
    - ColorID (Foreign Key, INT)
Relationships:
    - Many-to-One relationships with Brands, Sizes, and Colors.
Analysis:
    - The Products table stores details about the products available in the e-commerce platform.
    - ProductID serves as the primary key.
    - BrandID, SizeID, and ColorID establish relationships with other tables to categorize and describe products.
    - Price represents the product's price.

5. **Categories Table:**
Columns:
    - CategoryID (Primary Key, INT)
    - CategoryName (VARCHAR(255), NOT NULL)
    - Description (TEXT)
Relationships:
    - No direct relationships with other tables.
Analysis:
    - The Categories table defines product categories.
    - CategoryID is the primary key, and CategoryName stores the category's name.

6. **ProductCategories Table:**
Columns:
    - ProductCategoryID (Primary Key, INT)
    - ProductID (INT)
    - CategoryID (INT)
Relationships:
    - Many-to-Many relationships, i.e. Junction Table between Products and Categories.
Analysis:
    - The ProductCategories defines each product's categories (searchables)
    - ProductCategoryID is the primary key, ProductID and CategoryID will link the two tables.

7. **Brands Table:**
Columns:
    - BrandID (Primary Key, INT)
    - BrandName (VARCHAR(255), NOT NULL)
    - Description (TEXT)
Relationships:
    - No direct relationships with other tables.
Analysis:
    - The Brands table stores information about product brands.
    - BrandID is the primary key, and BrandName stores the brand's name.

8. **Sizes Table:**
Columns:
    - SizeID (Primary Key, INT)
    - SizeName (VARCHAR(20), NOT NULL)
    - Description (TEXT)
Relationships:
    - No direct relationships with other tables.
Analysis:
    - The Sizes table defines product sizes.
    - SizeID is the primary key, and SizeName stores the size's name.

9. **Colors Table:**
Columns:
    - ColorID (Primary Key, INT)
    - ColorName (VARCHAR(50), NOT NULL)
Relationships:
    - No direct relationships with other tables.
Analysis:
    - The Colors table defines product colors.
    - ColorID is the primary key, and ColorName stores the color's name.

10. **Stock Table:**
Columns:
    - StockID (Primary Key, INT)
    - ProductID (Foreign Key, INT)
    - Quantity (INT, NOT NULL)
Relationships:
    - Many-to-One relationship with Products (ProductID references Products.ProductID).
Analysis:
    - The Stock table is used to manage product stock levels.
    - StockID is the primary key, and ProductID links stock information to specific products.
