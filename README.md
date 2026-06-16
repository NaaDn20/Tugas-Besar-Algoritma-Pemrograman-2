# BangunIn - Sistem Informasi Supplier & Logistik 🏗️

**BangunIn** adalah aplikasi berbasis CLI (Command Line Interface) yang dirancang khusus untuk mengelola data master mitra supplier material konstruksi serta mencatat riwayat transaksi logistik secara terintegrasi. Proyek ini mengimplementasikan struktur data array of struct bertingkat dengan optimalisasi algoritma pencarian (searching) dan pengurutan (sorting) yang adaptif berdasarkan volume data.

---

## 📂 Struktur Repositori

* **flowchart/** : Folder berisi berkas visualisasi alur logika sistem
* **go.mod** : File manajemen modul Go
* **main.go** : Source code utama aplikasi BangunIn
* **README.md** : Dokumentasi proyek

---

## ✨ Fitur Utama Sistem

Aplikasi ini dilengkapi dengan 10 modul fungsionalitas utama yang dapat diakses secara dinamis melalui Menu Utama:

1. **Manajemen Master Supplier (CRUD)**
   * **Tampilkan Data:** Menyajikan profil lengkap supplier (ID, Kontak, Rating, Performa Transaksi) secara sekuensial.
   * **Tambah Data:** Registrasi supplier baru dengan validasi alokasi memori mencegah overflow (Maksimal 1000 data).
   * **Ubah Data:** Pembaruan informasi interaktif (fitur input `-` untuk mempertahankan data lama).
   * **Hapus Data:** Penghapusan data menggunakan metode pergeseran elemen array (shifting data).

2. **Sistem Manajemen Wilayah Bertingkat**
   * Pengelompokan lokasi supplier berdasarkan hierarki Provinsi ke Kabupaten/Kota.
   * Memungkinkan pembuatan wilayah baru secara dinamis saat registrasi supplier berlangsung dengan fitur auto-select.

3. **Pencatatan & Riwayat Transaksi**
   * Mengunci transaksi logistik baru ke ID Supplier tertentu yang valid.
   * Kalkulasi otomatis total harga belanja material.
   * Pelaporan fleksibel: Tampilkan seluruh log atau filter riwayat spesifik per ID Supplier.

4. **Algoritma Pencarian Adaptif (Searching)**
   * **Pencarian Nama:** Otomatis menggunakan Binary Search jika total data lebih dari 20 (didahului proses sorting alfabetis) dan menggunakan Sequential Search jika data kurang dari atau sama dengan 20.
   * **Pencarian Lokasi:** Menggunakan Sequential Search untuk menyisir dan menampilkan seluruh mitra yang berada di kawasan yang sama.

5. **Algoritma Pengurutan Efisien (Sorting)**
   * Menyusun peringkat supplier berdasarkan skor kepuasan pelanggan (rating).
   * Otomatis memicu fungsi Insertion Sort jika total data kurang dari atau sama dengan 20 demi efisiensi memori, dan beralih ke Selection Sort jika data lebih dari 20. Mendukung mode Ascending maupun Descending.

6. **Analitik & Statistik Data**
   * Kalkulasi rata-rata skor kepuasan (rating) seluruh supplier secara global.
   * Pemetaan kepadatan atau frekuensi jumlah supplier aktif di setiap titik wilayah operasional.

---

## 📊 Database Terinisialisasi (Data Dummy)

Untuk mempermudah pengujian seluruh kondisi logika algoritma, sistem ini sudah dilengkapi dengan fitur penyemaian data (data seeding) otomatis saat aplikasi pertama kali dijalankan:
* **21 Data Master Supplier** (Siap menguji kondisi algoritma pengurutan untuk data besar).
* **7 Data Transaksi Logistik Terbuku**.
* **2 Provinsi Utama** (DKI Jakarta & Jawa Timur) beserta daftar Kabupaten/Kota terkait.

---

## 🚀 Cara Menjalankan Program

### Prasyarat
Pastikan komputer Anda sudah terinstal Go (Golang).

### Langkah Eksekusi
1. Buka terminal di folder proyek ini.
2. Jalankan aplikasi langsung melalui terminal dengan perintah: **go run main.go**

---
*Proyek ini dikembangkan untuk memenuhi Tugas Besar mata kuliah Algoritma dan Pemrograman - Fakultas Informatika, Telkom University.*
