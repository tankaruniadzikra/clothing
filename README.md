# Cakra Clothing Store

## Background

Made as a project for Hacktiv8 Phase 1 Week 4 Pair Project

## Background Story

Sebagai seorang admin dari toko e-commerce yang fokus pada clothing industry, and diberikan sebuah aplikasi CLI untuk membantu anda menjalankan day-to-day operation. Di aplikasi CLI tersebut, anda dapat:

1. Mendaftarkan user atau customer baru
2. Login menggunakan credential anda

Setelah anda sukses login ke dalam aplikasi, anda dapat memilih sub menu selanjutnya, tergantung pada apa yang ingin anda lakukan:

1. Kelola product
2. Input order
3. Lihat report

### Kelola product

Di menu ini, anda dapat mengelola data product yang dijual oleh Cakra Clothing Store. Begitu anda memilih untuk `Kelola product` di menu sebelumnya, anda akan dapat akses untuk:

1. Lihat daftar product
2. Tambahkan product baru
3. Update product lama
4. Hapus product

Keempat menu itu cukup mudah untuk dimengerti, silahkan anda mencoba semua menu tersebut.

### Input order

Di menu ini, anda dapat mendaftarkan order baru dari customer. Cukup masuk ke menu ini, ikuti semua langkah di dalamnya, dan anda dapat dengan mudah mendaftarkan order baru.

### Lihat report

Di menu ini, anda dapat mengakses laporan-laporan yang berguna untuk keberlangsungan Cakra Clothing Store. Report yang tersedia untuk saat ini adalah:

1. Order overview
2. Order summary
3. Sales summary
4. Inventory summary
5. Product dengan pembelian terbanyak
6. User dengan pembelian terbanyak
7. Brand dengan pembelian terbanyak

Report tersebut disajikan dalam format table yang cukup user-friendly, jadi jangan takut untuk mencoba semua menu yang ada di sini.

Sekian penjelasan singkat dari aplikasi CLI Cakra Clothing Store. Have fun working as our admin team!

## Requirements

- go 1.19 or later
- External database dengan name: `cakra_clothing` untuk aplikasinya sendiri, dan nama: `cakra_clothing_test` untuk menjalankan test

## Running the Application

- Pastikan external database sudah hidup dan dapat diakses
- Jalankan aplikasinya dengan command ini:

```shell
    go run . -db_url="root:@tcp(localhost:3306)/cakra_clothing"
```

- Flag db_url dapat diganti dengan connection string ke database anda

## Team Member

- Nafisa Alfiani
- Tan Karunia Dzikra
