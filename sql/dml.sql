INSERT INTO Users (Email, Password, FirstName, LastName)
VALUES
('ica', 'ica', 'Nafisa', 'Alfiani'),
('dzikra', 'dzikra', 'Tan', 'Dzikra'),

INSERT INTO Brands (BrandName, Description)
VALUES
('Adidas', 'Brand sepatu dan pakaian olahraga terkemuka'),
('Nike', 'Brand sepatu dan pakaian atletik terkemuka'),
('Zara', 'Brand fesyen dengan desain terkini'),
('H&M', 'Brand pakaian terjangkau dan trendy');
('Puma', 'Brand terkenal dengan koleksi sepatu sporty'),
('Gucci', 'Brand fashion mewah dari Italia');

INSERT INTO Sizes (SizeName, Description)
VALUES
('S', 'Kecil'),
('M', 'Sedang'),
('L', 'Besar'),
('XL', 'Ekstra Besar');

INSERT INTO Colors (ColorName)
VALUES
('Merah'),
('Biru'),
('Hijau'),
('Kuning'),
('Hitam'),
('Putih'),
('Abu-abu'),

INSERT INTO Categories (CategoryName, Description)
VALUES
('Sepatu', 'Koleksi sepatu terbaru'),
('Pakaian', 'Koleksi pakaian terkini'),
('Aksesoris', 'Aksesoris fashion dan perhiasan'),
('Tas', 'Koleksi tas dari berbagai gaya dan ukuran'),
('Kacamata', 'Koleksi kacamata matahari dan bacaan'),
('Perhiasan', 'Perhiasan eksklusif untuk gaya Anda');

INSERT INTO Products (ProductName, Description, Price, Material, Weight, BrandID, SizeID, ColorID)
VALUES
('Sneaker MS92', 'Sepatu keren untuk gaya santai', 79.99, 'Suede', 0.6, 4, 3, 2),
('Tas Z23', 'Tas mewah dengan desain ikonik', 499.99, 'Kulit', 1.2, 6, 2, 3),
('Sepatu Olahraga D12', 'Sepatu olahraga unggulan', 99.99, 'Kulit sintetis', 0.75, 1, 2, 1),
('Kemeja Kasual ST67', 'Kemeja lengan panjang dengan desain modern', 39.99, 'Katun', 0.3, 3, 2, 3),
('Celana Pendek SH23', 'Celana pendek atletik', 29.99, 'Polyester', 0.2, 2, 1, 2);

INSERT INTO ProductCategories (ProductID, CategoryID)
VALUES
(1, 1),
(2, 4),
(3, 1),
(4, 2),
(5, 2);

INSERT INTO Stock (ProductID, Quantity)
VALUES
(1, 50),
(2, 75),
(3, 90),
(4, 25),
(5, 110);
