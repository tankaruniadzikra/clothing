-- 1. Users Table:
CREATE TABLE Users (
    UserID INT PRIMARY KEY,
    Email VARCHAR(100) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    FirstName VARCHAR(50),
    LastName VARCHAR(50)
);

-- 2. Orders Table:
CREATE TABLE Orders (
    OrderID INT PRIMARY KEY,
    UserID INT,
    OrderDate DATE NOT NULL,
    TotalAmount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- 3. OrderItems Table:
CREATE TABLE OrderItems (
    OrderItemID INT PRIMARY KEY,
    OrderID INT,
    ProductID INT,
    Quantity INT NOT NULL,
    PricePerUnit DECIMAL(10, 2) NOT NULL,
    TotalPrice DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

-- 4. Products Table:
CREATE TABLE Products (
    ProductID INT PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL,
    Description TEXT,
    Price DECIMAL(10, 2) NOT NULL,
    Material VARCHAR(255),
    Weight DECIMAL(8, 2),
    BrandID INT,
    SizeID INT,
    ColorID INT,
    FOREIGN KEY (BrandID) REFERENCES Brands(BrandID),
    FOREIGN KEY (SizeID) REFERENCES Sizes(SizeID),
    FOREIGN KEY (ColorID) REFERENCES Colors(ColorID)
);

-- 5. Categories Table:
CREATE TABLE Categories (
    CategoryID INT PRIMARY KEY,
    CategoryName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- 6. ProductCategories Table (Junction Table):
CREATE TABLE ProductCategories (
    ProductCategoryID INT PRIMARY KEY,
    ProductID INT,
    CategoryID INT,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

-- 7. Brands Table:
CREATE TABLE Brands (
    BrandID INT PRIMARY KEY,
    BrandName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- 8. Sizes Table:
CREATE TABLE Sizes (
    SizeID INT PRIMARY KEY,
    SizeName VARCHAR(20) NOT NULL,
    Description TEXT
);

-- 9. Colors Table:
CREATE TABLE Colors (
    ColorID INT PRIMARY KEY,
    ColorName VARCHAR(50) NOT NULL
);

-- 10. Stock Table:
CREATE TABLE Stock (
    StockID INT PRIMARY KEY,
    ProductID INT,
    Quantity INT NOT NULL,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);
