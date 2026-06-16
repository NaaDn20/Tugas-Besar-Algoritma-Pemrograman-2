# Tugas-Besar-Algoritma-Pemrograman-2

# BangunIn - Sistem Informasi Supplier & Logistik 🏗️

**BangunIn** adalah aplikasi yang dirancang khusus untuk mengelola data master mitra supplier material konstruksi serta mencatat riwayat transaksi logistik secara terintegrasi. Proyek ini mengimplementasikan struktur data array of struct bertingkat dengan optimalisasi algoritma pencarian (*searching*) dan pengurutan (*sorting*) yang adaptif berdasarkan volume data.

---

## 📂 Struktur Repositori

```text
├── flowchart/           # Folder berisi berkas visualisasi alur logika sistem
├── go.mod               # File manajemen modul Go
├── main.go              # Source code utama aplikasi BangunIn
└── README.md            # Dokumentasi proyek

```✨ Fitur Utama SistemAplikasi ini dilengkapi dengan 10 modul fungsionalitas utama yang dapat diakses secara dinamis melalui Menu Utama:Manajemen Master Supplier (CRUD)Tampilkan Data: Menyajikan profil lengkap supplier (ID, Kontak, Rating, Performa Transaksi) secara sekuensial.Tambah Data: Registrasi supplier baru dengan validasi alokasi memori mencegah overflow ($MAX = 1000$).Ubah Data: Pembaruan informasi interaktif (fitur input - untuk mempertahankan data lama).Hapus Data: Penghapusan data menggunakan metode pergeseran elemen array (shifting data).Sistem Manajemen Wilayah BertingkatPengelompokan lokasi supplier berdasarkan hierarki Provinsi ➡️ Kabupaten/Kota.Memungkinkan pembuatan wilayah baru secara dinamis saat registrasi supplier berlangsung dengan fitur auto-select.Pencatatan & Riwayat TransaksiMengunci transaksi logistik baru ke ID Supplier tertentu yang valid.Kalkulasi otomatis total harga belanja material.Pelaporan fleksibel: Tampilkan seluruh log atau filter riwayat spesifik per ID Supplier.Algoritma Pencarian Adaptif (Searching)Pencarian Nama: Otomatis menggunakan Binary Search jika total data $> 20$ (didahului proses sorting alfabetis) dan menggunakan Sequential Search jika data $\le 20$.Pencarian Lokasi: Menggunakan Sequential Search untuk menyisir dan menampilkan seluruh mitra yang berada di kawasan yang sama.Algoritma Pengurutan Efisien (Sorting)Menyusun peringkat supplier berdasarkan skor kepuasan pelanggan (rating).Otomatis memicu fungsi Insertion Sort jika total data $\le 20$ demi efisiensi memori, dan beralih ke Selection Sort jika total data $> 20$. Supports mode Ascending maupun Descending.Analitik & Statistik DataKalkulasi rata-rata skor kepuasan (rating) seluruh supplier secara global.Pemetaan kepadatan atau frekuensi jumlah supplier aktif di setiap titik wilayah operasional.📊 Database Terinisialisasi (Data Dummy)Untuk mempermudah pengujian seluruh kondisi logika algoritma (sorting & searching ekstrem), sistem ini sudah dilengkapi dengan fitur penyemaian data (data seeding) otomatis saat aplikasi pertama kali dijalankan:21 Data Master Supplier (Siap menguji kondisi algoritma pengurutan $> 20$ data).7 Data Transaksi Logistik Terbuku.2 Provinsi Utama (DKI Jakarta & Jawa Timur) beserta daftar Kabupaten/Kota terkait.🚀 Cara Menjalankan ProgramPrasyaratPastikan komputer Anda sudah terinstal Go (Golang) versi terbaru.Langkah EksekusiClone Repositori ini ke laptop Anda:Bashgit clone [https://github.com/USERNAME_KAMU/NAMA_REPO.git](https://github.com/USERNAME_KAMU/NAMA_REPO.git)
cd NAMA_REPO
Jalankan aplikasi langsung melalui terminal:Bashgo run main.go
Proyek ini dikembangkan untuk memenuhi Tugas Besar mata kuliah Algoritma dan Pemrograman - Fakultas Informatika, Telkom University.
