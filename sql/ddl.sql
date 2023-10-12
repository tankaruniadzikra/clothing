DROP DATABASE IF EXISTS cakra_clothing;

CREATE DATABASE cakra_clothing;

USE cakra_clothing;

-- Users Table:
CREATE TABLE Users (
    UserID INT PRIMARY KEY AUTO_INCREMENT,
    Email VARCHAR(100) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    FirstName VARCHAR(50),
    LastName VARCHAR(50)
);

-- Brands Table:
CREATE TABLE Brands (
    BrandID INT PRIMARY KEY AUTO_INCREMENT,
    BrandName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- Sizes Table:
CREATE TABLE Sizes (
    SizeID INT PRIMARY KEY AUTO_INCREMENT,
    SizeName VARCHAR(20) NOT NULL,
    Description TEXT
);

-- Colors Table:
CREATE TABLE Colors (
    ColorID INT PRIMARY KEY AUTO_INCREMENT,
    ColorName VARCHAR(50) NOT NULL
);

-- Categories Table:
CREATE TABLE Categories (
    CategoryID INT PRIMARY KEY AUTO_INCREMENT,
    CategoryName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- Products Table:
CREATE TABLE Products (
    ProductID INT PRIMARY KEY AUTO_INCREMENT,
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

-- ProductCategories Table (Junction Table):
CREATE TABLE ProductCategories (
    ProductCategoryID INT PRIMARY KEY AUTO_INCREMENT,
    ProductID INT,
    CategoryID INT,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID) ON DELETE CASCADE,
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

-- Stock Table:
CREATE TABLE Stock (
    StockID INT PRIMARY KEY AUTO_INCREMENT,
    ProductID INT,
    Quantity INT NOT NULL,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID) ON DELETE CASCADE
);

-- Orders Table:
CREATE TABLE Orders (
    OrderID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT,
    OrderDate DATE NOT NULL,
    TotalAmount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- OrderItems Table:
CREATE TABLE OrderItems (
    OrderItemID INT PRIMARY KEY AUTO_INCREMENT,
    OrderID INT,
    ProductID INT,
    Quantity INT NOT NULL,
    PricePerUnit DECIMAL(10, 2) NOT NULL,
    TotalPrice DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID) ON DELETE CASCADE,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID) ON DELETE CASCADE
);
