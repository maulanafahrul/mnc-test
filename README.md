# mnc-test

Proyek ini adalah contoh implementasi manajemen data menggunakan framework web Gin-Gonic di Golang. Proyek ini mencakup beberapa fitur seperti manajemen pelanggan, pembayaran, pengguna, dan otentikasi.

## Instalasi

1. Pastikan Anda memiliki Go (Golang) diinstal. Jika belum, Anda dapat mengunduhnya dari [sini](https://golang.org/dl/).

2. Clone repositori ini:
   (https://github.com/maulanafahrul/mnc-test.git)
   
## Penggunaan

### Autentikasi

Proyek ini mendukung autentikasi dengan endpoint /login untuk masuk dan /logout untuk keluar. Setelah masuk, Anda akan mendapatkan token JWT yang digunakan untuk mengakses endpoint lainnya.

## Manajemen Pengguna

GET /user/:username: Mendapatkan informasi pengguna berdasarkan nama pengguna.

POST /user: Menambahkan pengguna baru.

PUT /user: Mengupdate informasi pengguna.

DELETE /user/:username: Menghapus pengguna berdasarkan nama pengguna.

## Manajemen Pelanggan

GET /customer/:fullname: Mendapatkan informasi pelanggan berdasarkan nama lengkap.

POST /customer: Menambahkan pelanggan baru.

PUT /customer: Mengupdate informasi pelanggan.

DELETE /customer/:fullname: Menghapus pelanggan berdasarkan nama lengkap.

## Manajemen Pembayaran

GET /payment/:id: Mendapatkan informasi pembayaran berdasarkan ID.

POST /payment: Menambahkan pembayaran baru.
