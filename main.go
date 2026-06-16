package main
import "fmt"

const MAX int = 1000

type kontak struct{
	nomor string
	email string
}

type supplier struct{
	id int
	nama string
	lokasi string
	material string
	detailKontak kontak
	jumlahTransaksi int
	rating float64
}

type transaksi struct {
    idTransaksi int
	tanggal string
    idSupplier  int
    namaSupplier string
    material    string
    jumlah      int
    totalHarga  float64
}

type lokasi struct{
	provinsi string
	jumlahKabkot int
	kabkot [MAX]string
}

var daftar [MAX]supplier
var daftarTransaksi [MAX]transaksi
var dataLok [MAX]lokasi

func isiDataLokasi(jumlahProv *int) {
	dataLok[0].provinsi = "DKI_Jakarta"
	dataLok[0].kabkot[0] = "Jakarta_Pusat"
	dataLok[0].kabkot[1] = "Jakarta_Utara"
	dataLok[0].kabkot[2] = "Jakarta_Barat"
	dataLok[0].kabkot[3] = "Jakarta_Selatan"
	dataLok[0].kabkot[4] = "Jakarta_Timur"
	dataLok[0].kabkot[5] = "Kepulauan_Seribu"
	dataLok[0].jumlahKabkot = 6

	dataLok[1].provinsi = "Jawa_Timur"
	dataLok[1].kabkot[0] = "Surabaya"
	dataLok[1].kabkot[1] = "Malang"
	dataLok[1].kabkot[2] = "Sidoarjo"
	dataLok[1].kabkot[3] = "Gresik"
	dataLok[1].kabkot[4] = "Mojokerto"
	dataLok[1].kabkot[5] = "Pasuruan"
	dataLok[1].kabkot[6] = "Kediri"
	dataLok[1].kabkot[7] = "Blitar"
	dataLok[1].kabkot[8] = "Madiun"
	dataLok[1].kabkot[9] = "Jember"
	dataLok[1].jumlahKabkot = 10

	*jumlahProv = 2
}

func isiDummySup(jumlahSupplier *int, nextID *int){
	daftar[0] = supplier{1, "TokoBangunanMaju", "Jakarta_Pusat", "Semen", kontak{"081234567890", "majubangunan@gmail.com"}, 0, 8.5}
	daftar[1] = supplier{2, "SumberMakmur", "Jakarta_Pusat", "Bata_Merah", kontak{"082345678901", "sumbermakmur@gmail.com"}, 0, 7.8}
	daftar[2] = supplier{3, "KaryaAbadi", "Jakarta_Barat", "Pasir", kontak{"083456789012", "karyaabadi@gmail.com"}, 0, 9.2}
	daftar[3] = supplier{4, "BerkahJaya", "Jakarta_Barat", "Besi_Beton", kontak{"084567890123", "berkahjaya@gmail.com"}, 0, 8.9}
	daftar[4] = supplier{5, "MandiriMaterial", "Surabaya", "Kayu", kontak{"085678901234", "mandirimaterial@gmail.com"}, 0, 6.4}
	daftar[5] = supplier{6, "SejahteraBangun", "Surabaya", "Semen", kontak{"086789012345", "sejahterabangun@gmail.com"}, 0, 9.5}
	daftar[6] = supplier{7, "PrimaKonstruksi", "Malang", "Genteng", kontak{"087890123456", "primakonstruksi@gmail.com"}, 0, 7.3}
	daftar[7] = supplier{8, "JayaAbadi", "Sidoarjo", "Besi_Beton", kontak{"088901234567", "jayaabadi@gmail.com"}, 0, 8.1}
	daftar[8] = supplier{9, "MitraBangunan", "Sidoarjo", "Pasir", kontak{"089012345678", "mitrabangunan@gmail.com"}, 0, 9.7}
	daftar[9] = supplier{10, "UsahaTama", "Kediri", "Kayu", kontak{"081122334455", "usahatama@gmail.com"}, 0, 5.9}
	*jumlahSupplier = 10
	*nextID = 11

	daftar[10] = supplier{11, "BumiPerkasa", "Jakarta_Pusat", "Bata_Merah", kontak{"081233445566", "bumiperkasa@gmail.com"}, 0, 8.3}
	daftar[11] = supplier{12, "SuryaMaterial", "Jakarta_Utara", "Pasir", kontak{"082344556677", "suryamaterial@gmail.com"}, 0, 6.1}
	daftar[12] = supplier{13, "AgungJaya", "Jakarta_Selatan", "Semen", kontak{"083455667788", "agungjaya@gmail.com"}, 0, 9.0}
	daftar[13] = supplier{14, "CahayaBangun", "Malang", "Besi_Beton", kontak{"084566778899", "cahayabangun@gmail.com"}, 0, 7.6}
	daftar[14] = supplier{15, "DeltaKonstruksi", "Surabaya", "Kayu", kontak{"085677889900", "deltakonstruksi@gmail.com"}, 0, 7.1}
	daftar[15] = supplier{16, "EmasAbadi", "Gresik", "Genteng", kontak{"086788990011", "emasabadi@gmail.com"}, 0, 8.6}
	daftar[16] = supplier{17, "FortunaBangunan", "Jakarta_Barat", "Pasir", kontak{"087899001122", "fortunabangunan@gmail.com"}, 0, 8.0}
	daftar[17] = supplier{18, "GrahaMakmur", "Sidoarjo", "Semen", kontak{"088900112233", "grahamakmur@gmail.com"}, 0, 6.8}
	daftar[18] = supplier{19, "HarapanJaya", "Malang", "Bata_Merah", kontak{"089011223344", "harapanjaya@gmail.com"}, 0, 9.3}
	daftar[19] = supplier{20, "IndahMaterial", "Surabaya", "Besi_Beton", kontak{"081344556677", "indahmaterial@gmail.com"}, 0, 7.4}
	daftar[20] = supplier{21, "JempolBangun", "Gresik", "Kayu", kontak{"082455667788", "jempolbangun@gmail.com"}, 0, 8.8}
	*jumlahSupplier = 21
	*nextID = 22
}

func isiDummyTransaksi(jumlahTransaksi *int, nextIDTransaksi *int){
	daftarTransaksi[0] = transaksi{1, "01/01/2026", 1, "Toko Bangunan Maju", "Semen", 50, 3250000}
	daftarTransaksi[1] = transaksi{2, "02/03/2026", 6, "Sejahtera Bangun", "Semen", 100, 6400000}
	daftarTransaksi[2] = transaksi{3, "05/03/2026", 3, "Karya Abadi", "Pasir", 30, 1500000}
	daftarTransaksi[3] = transaksi{4, "20/03/2026", 1, "Toko Bangunan Maju", "Semen", 75, 4875000}
	daftarTransaksi[4] = transaksi{5, "11/04/2026", 6, "Sejahtera Bangun", "Semen", 60, 3840000}
	daftarTransaksi[5] = transaksi{6, "30/04/2026", 1, "Toko Bangunan Maju", "Semen", 40, 2600000}
	daftarTransaksi[6] = transaksi{7, "01/05/2026", 8, "Jaya Abadi", "Besi_Beton", 20, 4200000}

	daftar[0].jumlahTransaksi = 3
	daftar[2].jumlahTransaksi = 1
	daftar[5].jumlahTransaksi = 2
	daftar[7].jumlahTransaksi = 1

	*jumlahTransaksi = 7
	*nextIDTransaksi = 8
}

func aturAtauPilihLokasi(jumlahProv *int, mode string) string{
	var lokPilih string = ""
	idxProv := -1
	selesai := false

	for !selesai{
		if idxProv == -1{
			fmt.Println("\n	=========== Data Lokasi ===========")
			fmt.Println("	Daftar Provinsi:")
			for i := 0; i < *jumlahProv; i++{
				fmt.Printf("	%d. %s\n", i+1, dataLok[i].provinsi)
			}

			batasPil := *jumlahProv
			if mode == "tambah" || mode == "atur" || mode == "ubah"{
				batasPil = batasPil + 1
				if mode == "tambah" || mode == "ubah"{
					fmt.Printf("	%d. Lainnya\n", *jumlahProv+1)
				}else{
					fmt.Printf("	%d. Tambah Provinsi Baru\n", *jumlahProv+1)
				}
			}
			fmt.Println("	0. Batal")
			fmt.Print("	Pilihan Anda: ")

			var pilProv int = -1
			for pilProv < 0 || pilProv > batasPil{
				fmt.Scan(&pilProv)
				if pilProv < 0 || pilProv > batasPil{
					fmt.Println("	Pilihan tidak valid.")
					fmt.Print("	Pilihan Anda: ")
				}
			}

			if pilProv == 0{
				selesai = true
			}else if (mode == "tambah" || mode == "atur" || mode == "ubah" ) && pilProv == batasPil{
				fmt.Print("	Masukkan nama provinsi baru: ")
				var provBaru string
				fmt.Scan(&provBaru)
				dataLok[*jumlahProv].provinsi = provBaru
				dataLok[*jumlahProv].jumlahKabkot = 0

				fmt.Printf("	Provinsi %s berhasil ditambahkan.\n", provBaru)

				if mode == "tambah" || mode == "ubah"{
					idxProv = *jumlahProv
				}

				*jumlahProv++
			}else{
				idxProv = pilProv - 1
			}
		}

		if idxProv != -1 && !selesai{
			selesaiKabkot := false
			for !selesaiKabkot{
				fmt.Printf("\n	== Kab/Kota di %s ==\n", dataLok[idxProv].provinsi)
				for i := 0; i < dataLok[idxProv].jumlahKabkot; i++{
					fmt.Printf("	%d. %s\n", i+1, dataLok[idxProv].kabkot[i])
				}

				batasKabkot := dataLok[idxProv].jumlahKabkot
				if mode == "tambah" || mode == "atur" || mode == "ubah"{
					batasKabkot = batasKabkot + 1
					if mode == "tambah" || mode == "ubah"{
						fmt.Printf("	%d. Lainnya\n", dataLok[idxProv].jumlahKabkot+1)
					}else{
						fmt.Printf("	%d. Tambahkan Kab/Kota Baru\n", dataLok[idxProv].jumlahKabkot+1)
					}
				}

				fmt.Println("	0. Kembali")
				fmt.Print("	Pilihan Anda: ")

				var pilKabkot int = -1
				for pilKabkot < 0 || pilKabkot > batasKabkot{
					fmt.Scan(&pilKabkot)
					if pilKabkot < 0 || pilKabkot > batasKabkot{
						fmt.Println("	Pilihan tidak valid.")
						fmt.Print("	Pilihan Anda: ")
					}
				}

				if pilKabkot == 0{
					idxProv = -1
					selesaiKabkot = true
				}else if (mode == "tambah" || mode == "atur" || mode == "ubah") && pilKabkot == batasKabkot{
					fmt.Print("	Masukkan nama kab/kota baru: ")
					var kabkotBaru string
					fmt.Scan(&kabkotBaru)
					dataLok[idxProv].kabkot[dataLok[idxProv].jumlahKabkot] = kabkotBaru
					dataLok[idxProv].jumlahKabkot++

					fmt.Printf("	Kab/Kota %s berhasil ditambahkan.\n", kabkotBaru)

					if mode == "tambah" || mode == "ubah"{
						lokPilih = kabkotBaru
						selesaiKabkot = true
						selesai = true
					}
				}else{
					if mode == "tambah" || mode == "ubah" || mode == "cari"{
						lokPilih = dataLok[idxProv].kabkot[pilKabkot-1]
						selesaiKabkot = true
						selesai = true
					}else{
						fmt.Printf("	Kab/Kota %s sudah terdaftar di %s.\n", dataLok[idxProv].kabkot[pilKabkot-1], dataLok[idxProv].provinsi)
					}
				}
			}
		}
	}

	return lokPilih
}

func tampilkanSemuaSupplier(jumlahSupplier int){
	if jumlahSupplier == 0 {
		fmt.Println("Belum ada data supplier.")
	}else{
		fmt.Println("\n========== Daftar Semua Supplier ===========")
		for i := 0; i < jumlahSupplier; i++{
			tampilkanSupplier(daftar[i])
			fmt.Println("============================================")
		}
	}
}

func tampilkanSupplier(sup supplier){
	fmt.Println("ID 			:", sup.id)
	fmt.Println("Nama 			:", sup.nama)
	fmt.Println("Lokasi			:", sup.lokasi)
	fmt.Println("Material 		:", sup.material)
	fmt.Println("No HP 			:", sup.detailKontak.nomor)
	fmt.Println("E-mail			:", sup.detailKontak.email)
	fmt.Println("Total Transaksi		:", sup.jumlahTransaksi)
	fmt.Println("Rating	 		:", sup.rating)
}

func tambahSupplier(jumlahSupplier, nextID *int, jumlahProv *int){
	if *jumlahSupplier >= MAX{
		fmt.Println("Database supplier penuh.")
	}else{
		daftar[*jumlahSupplier].id = *nextID

		fmt.Println("Masukkan data supplier baru")

		fmt.Print("Nama Supplier: ")
		fmt.Scan(&daftar[*jumlahSupplier].nama)

		fmt.Print("Pilih Lokasi Supplier: ")
		mode := "tambah"	
		daftar[*jumlahSupplier].lokasi = aturAtauPilihLokasi(jumlahProv, mode)
		if daftar[*jumlahSupplier].lokasi == ""{
			fmt.Println("Penambahan supplier dibatalkan.")
		}else{
			fmt.Print("Jenis Material: ")
			fmt.Scan(&daftar[*jumlahSupplier].material)
	
			fmt.Print("Nomor HP: ")
			fmt.Scan(&daftar[*jumlahSupplier].detailKontak.nomor)
	
			fmt.Print("E-mail: ")
			fmt.Scan(&daftar[*jumlahSupplier].detailKontak.email)
	
			fmt.Print("Rating (/10): ")
			fmt.Scan(&daftar[*jumlahSupplier].rating)

			daftar[*jumlahSupplier].jumlahTransaksi = 0
	
			*jumlahSupplier++
			*nextID++
	
			fmt.Println("Supplier berhasil ditambahkan.")
		}
	}
}

func ubahSupplier(jumlahSupplier int, jumlahProv *int){
	if jumlahSupplier == 0 {
		fmt.Println("Belum ada data supplier.")
	}else{
		var idCari int
		fmt.Print("\nMasukkan ID Supplier yang ingin diubah: ")
		fmt.Scan(&idCari)

		idx := -1
		ketemu := false
		for i := 0; i < jumlahSupplier && !ketemu; i++{
			if daftar[i].id == idCari{
				idx = i
				ketemu = true
			}
		}

		if idx == -1{
			fmt.Println("Supplier dengan ID tersebut tidak ditemukan.")
		}else{
			fmt.Printf("\n=== Ubah Supplier: %s (ID: %d) ===\n", daftar[idx].nama, daftar[idx].id)
			fmt.Println("(Ketik '-' untuk tidak mengubah data tersebut)")

			var input string

			fmt.Printf("Nama Supplier [%s]: ", daftar[idx].nama)
			fmt.Scan(&input)
			if input != "-" {
				daftar[idx].nama = input
			}

			fmt.Printf("Lokasi Supplier [%s] (ketik 'y' untuk mengubah): ", daftar[idx].lokasi)
			fmt.Scan(&input)
			if input == "y" {
				mode := "ubah"
				lokasiBaru := aturAtauPilihLokasi(jumlahProv, mode)
				if lokasiBaru != ""{
					daftar[idx].lokasi = lokasiBaru
				}
			}

			fmt.Printf("Jenis Material Supplier [%s]: ", daftar[idx].material)
			fmt.Scan(&input)
			if input != "-" {
				daftar[idx].material = input
			}

			fmt.Printf("Nomor Telepon [%s]: ", daftar[idx].detailKontak.nomor)
			fmt.Scan(&input)
			if input != "-" {
				daftar[idx].detailKontak.nomor = input
			}

			fmt.Printf("E-mail [%s]: ", daftar[idx].detailKontak.email)
			fmt.Scan(&input)
			if input != "-" {
				daftar[idx].detailKontak.email = input
			}

			var RatingBaru float64
			fmt.Printf("Rating [%.1f (/10)] (Ketik -1 untuk tidak mengubah): ", daftar[idx].rating)
			fmt.Scan(&RatingBaru)
			if RatingBaru >= 0 && RatingBaru <= 10 {
				daftar[idx].rating = RatingBaru
			}

			fmt.Println("Data supplier berhasil diubah.")
		}
	}
}

func hapusSupplier(jumlahSupplier *int){
	if *jumlahSupplier == 0 {
		fmt.Println("Belum ada data supplier.")
	}else{
		var id int
		var idx int = -1
		fmt.Print("Masukkan ID supplier yang akan dihapus: ")
		fmt.Scan(&id)
		for i:= 0; i < *jumlahSupplier; i++ {
			if daftar[i].id == id {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("Supplier tidak ditemukan.")
		}else{
			for i:= idx; i < *jumlahSupplier-1; i++ {
				daftar[i] = daftar[i+1]
			}
			*jumlahSupplier--
			fmt.Println("Data supplier berhasil dihapus.")
		}
	}
}

func sequentialSearch(keyword string, byNama bool, jumlahSupplier int) ([MAX]int, int){
	var daftarIndex [MAX]int
	jumlahKetemu := 0

	for i := 0; i < jumlahSupplier; i++{
		sama := false
		if byNama{
			if daftar[i].nama == keyword{
				sama = true
			}
		}else{
			if daftar[i].lokasi == keyword{
				sama = true
			}
		}

		if sama {
			daftarIndex[jumlahKetemu] = i
			jumlahKetemu++
		}
	}

	return daftarIndex, jumlahKetemu
}

func urutNama(jumlahSupplier int) [MAX]supplier{
	var temp [MAX]supplier
	for i := 0; i < jumlahSupplier; i++{
		temp[i] = daftar[i]
	}

	for i := 0; i < jumlahSupplier - 1; i++{
		min := i
		for j := i + 1; j < jumlahSupplier; j++{
			if temp[min].nama > temp[j].nama{
				min = j
			}
		}
		temps := temp[i]
		temp[i] = temp[min]
		temp[min] = temps
	}

	return temp
}

func binarySearch(data [MAX]supplier, keyword string, jumlahSupplier int) int{
	found := -1
	kiri := 0
	kanan := jumlahSupplier-1

	for kiri <= kanan && found == -1{
		mid := (kiri + kanan) / 2
		if keyword < data[mid].nama {
			kanan = mid - 1
		}else if keyword > data[mid].nama {
			kiri = mid + 1
		}else{
			found = mid
		} 
	}
	return found
}

func cariSupplier(jumlahSupplier int, jumlahProv int){
	if jumlahSupplier == 0 {
		fmt.Println("Belum ada data supplier.")
	}else{
		fmt.Println("\n============ Cari Data Supplier =============")
		var pilCari int = -1
		for pilCari < 0 || pilCari >2{
			fmt.Println("Cari Berdasarkan:")
			fmt.Println("	1. Nama Supplier")
			fmt.Println("	2. Lokasi Supplier")
			fmt.Println("	0. Kembali")
			fmt.Print("Pilihan Anda (0-2): ")
			fmt.Scan(&pilCari)

			if pilCari < 0 || pilCari >2{
				fmt.Println("\nPilihan tidak valid.")
				fmt.Println()
			}
		}

		if pilCari != 0{
			byNama := pilCari == 1

			var keyword string
			if byNama {
				fmt.Print("Masukkan nama supplier yang ingin dicari: ")
				fmt.Scan(&keyword)
			}else{
				fmt.Print("Pilih lokasi supplier yang ingin dicari: ")
				mode := "cari"
				keyword = aturAtauPilihLokasi(&jumlahProv, mode)
			}
			
			if jumlahSupplier <= 20 || !byNama {
				daftarIndex, jumlahKetemu := sequentialSearch(keyword, byNama, jumlahSupplier)
				if jumlahKetemu > 0{
					fmt.Println("\nSupplier ditemukan!")
					fmt.Println("============================================")
					for i := 0; i < jumlahKetemu; i++{
						idx := daftarIndex[i]
						tampilkanSupplier(daftar[idx])
						fmt.Println("============================================")
					}
					if !byNama{
						fmt.Printf("Jumlah Supplier Ditemukan: %d\n", jumlahKetemu)
					}
				}else{
					fmt.Println("\nSupplier tidak ditemukan.")
				}
			}else{
				terurut := urutNama(jumlahSupplier)
				idx := binarySearch(terurut, keyword, jumlahSupplier)
				if idx != -1{
					fmt.Println("\nSupplier ditemukan!")
					fmt.Println("============================================")
					tampilkanSupplier(terurut[idx])
					fmt.Println("============================================")
				}else{
					fmt.Println("\nSupplier tidak ditemukan.")
				}
			}
		}
	}
}

func selectionSort(asc bool, jumlahSupplier int) {
	for i := 0; i < jumlahSupplier - 1; i++{
		idx := i
		for j := i + 1; j < jumlahSupplier; j++{
			if asc{
				if daftar[idx].rating > daftar[j].rating{
					idx = j
				}
			}else{
				if daftar[idx].rating < daftar[j].rating{
					idx = j
				}
			}
		}
		temp := daftar[i]
		daftar[i] = daftar[idx]
		daftar[idx] = temp
	}
}

func insertionSort(asc bool, jumlahSupplier int){
	for i := 1; i <= jumlahSupplier - 1; i++{
		j := i
		temp := daftar[j]
		if asc{
			for j > 0 && temp.rating < daftar[j-1].rating{
				daftar[j] = daftar[j-1]
				j = j - 1
			}
		}else{
			for j > 0 && temp.rating > daftar[j-1].rating{
				daftar[j] = daftar[j-1]
				j = j - 1
			}
		}
		daftar[j] = temp
	}
}

func urutkanSupplier(jumlahSupplier int){
	if jumlahSupplier == 0 {
		fmt.Println("Belum ada data supplier.")
	}else{
		fmt.Println("\n== Urut Data Supplier Berdasarkan Rating ==")
		var pilUrut int = -1
		for pilUrut < 0 || pilUrut >2{
			fmt.Println("Urut Secara:")
			fmt.Println("	1. Terendah ke Tertinggi")
			fmt.Println("	2. Tertinggi ke Terendah")
			fmt.Println("	0. Kembali")
			fmt.Print("Pilihan Anda (0-2): ")
			fmt.Scan(&pilUrut)

			if pilUrut < 0 || pilUrut >2{
				fmt.Println("\nPilihan tidak valid.")
				fmt.Println()
			}
		}

		if pilUrut != 0{
			asc := pilUrut == 1

			if jumlahSupplier <= 20{
				insertionSort(asc, jumlahSupplier)
			}else{
				selectionSort(asc, jumlahSupplier)
			}

			if asc {
				fmt.Println("\nDaftar Supplier (Rating Terendah -> Tertinggi)")
			}else{
				fmt.Println("\nDaftar Supplier (Rating Tertinggi -> Terendah)")
			}

			fmt.Println("============================================")
			for i := 0; i < jumlahSupplier; i++{
				tampilkanSupplier(daftar[i])
				fmt.Println("============================================")
			}
		}
	}
}

func tampilkanStatistik(jumlahSupplier int){
	if jumlahSupplier == 0 {
		fmt.Println("Belum ada data supplier.")
	}else{
		var totalRating float64
		var wilayah [MAX]string
		var jumlahWilayah [MAX]int
		var banyakWilayah int = 0

		for i := 0; i < jumlahSupplier; i++{
			totalRating += daftar[i].rating
		}

		for i := 0; i < jumlahSupplier; i++{
			idx := -1
			for j := 0; j < banyakWilayah; j++{
				if wilayah[j] == daftar[i].lokasi {
					idx = j
				}
			}

			if idx == -1{
				wilayah[banyakWilayah] = daftar[i].lokasi
				jumlahWilayah[banyakWilayah] = 1
				banyakWilayah++
			}else{
				jumlahWilayah[idx]++
			}
		}

		rataRata := totalRating / float64(jumlahSupplier)

		fmt.Println("\n================ STATISTIK ================")
		fmt.Println("Jumlah Supplier :", jumlahSupplier)
		fmt.Printf("Rata-rata Skor Kepuasan : %.2f\n", rataRata)

		fmt.Println("\nJumlah Supplier per Wilayah: ")

		for i := 0; i < banyakWilayah; i++{
			fmt.Printf("%s : %d supplier\n", wilayah[i], jumlahWilayah[i])
		}
		fmt.Println("============================================")
	}
}

func catatTransaksi(jumlahSupplier int, jumlahTransaksi *int, nextIDTransaksi *int){
	var tanggal string
	if jumlahSupplier == 0{
		fmt.Println("Belum ada supplier.")
	}else{
		var idSup int
		fmt.Print("Masukkan ID Supplier: ")
		fmt.Scan(&idSup)

		idx := -1
		for i := 0; i < jumlahSupplier; i++{
			if daftar[i].id == idSup{
				idx = i
			}
		}

		if idx == -1{
			fmt.Println("Supplier tidak ditemukan.")
		}else{
			var jumlahBarang int
			var hargaSatuan float64
			fmt.Print("Tanggal (DD/MM/YYYY): ")
			fmt.Scan(&tanggal)

			fmt.Print("Jumlah Barang: ")
			fmt.Scan(&jumlahBarang)

			fmt.Print("Harga Satuan: ")
			fmt.Scan(&hargaSatuan)

			daftarTransaksi[*jumlahTransaksi].idTransaksi = *nextIDTransaksi
			daftarTransaksi[*jumlahTransaksi].tanggal = tanggal
			daftarTransaksi[*jumlahTransaksi].idSupplier = daftar[idx].id
			daftarTransaksi[*jumlahTransaksi].namaSupplier = daftar[idx].nama
			daftarTransaksi[*jumlahTransaksi].material = daftar[idx].material
			daftarTransaksi[*jumlahTransaksi].jumlah = jumlahBarang
			daftarTransaksi[*jumlahTransaksi].totalHarga = float64(jumlahBarang) * hargaSatuan

			*jumlahTransaksi++
			*nextIDTransaksi++
			daftar[idx].jumlahTransaksi++

			fmt.Println("Transaksi berhasil dicatat.")
		}
	}
}

func tampilkanDetailTransaksi(data transaksi){
	fmt.Printf("ID Transaksi	: %d\n", data.idTransaksi)
	fmt.Printf("Tanggal		: %s\n", data.tanggal)
	fmt.Printf("ID Supplier	: %d\n", data.idSupplier)
	fmt.Printf("Nama Supplier	: %s\n", data.namaSupplier)
	fmt.Printf("Jenis Material	: %s\n", data.material)
	fmt.Printf("Jumlah Barang	: %d\n", data.jumlah)
	fmt.Printf("Total Harga	: Rp.%.0f\n", data.totalHarga)
}

func tampilkanRiwayatTransaksi(jumlahTransaksi int){
	if jumlahTransaksi == 0{
		fmt.Println("Belum ada riwayat transaksi.")
	}else{
		fmt.Println("\n============== Riwayat Transaksi ===============")
		var pil int = -1
		for pil < 0 || pil >2{
			fmt.Println("Opsi:")
			fmt.Println("	1. Tampilkan Seluruh Riwayat Transaksi")
			fmt.Println("	2. Tampilkan Transaksi Per Supplier")
			fmt.Println("	0. Kembali")
			fmt.Print("Pilihan Anda (0-2): ")
			fmt.Scan(&pil)

			if pil < 0 || pil >2{
				fmt.Println("\nPilihan tidak valid.")
				fmt.Println()
			}
		}

		if pil == 1{
			fmt.Println("=========== Semua Riwayat Transaksi ============")
			for i := 0; i < jumlahTransaksi; i++{
				tampilkanDetailTransaksi(daftarTransaksi[i])
				fmt.Println("============================================")
			}
			fmt.Printf("Total Transaksi: %d\n", jumlahTransaksi)
		}else if pil == 2{
			var idSup int

			fmt.Print("Masukkan ID Supplier yang ingin dicari: ")
			fmt.Scan(&idSup)

			fmt.Printf("\n====== Riwayat Transaksi Supplier ID %d =======\n", idSup)

			ketemu := false
			banyak := 1
			for i := 0; i < jumlahTransaksi; i++{
				if daftarTransaksi[i].idSupplier == idSup{
					fmt.Printf("Transaksi ke-%d\n", banyak)
					tampilkanDetailTransaksi(daftarTransaksi[i])
					ketemu = true
					banyak++
					fmt.Println("============================================")
				}
			}

			if !ketemu {
				fmt.Println("Tidak ada data transaksi untuk ID Supplier tersebut.")
				fmt.Println("============================================")
			}
		}
	}
}

func main(){
	var jumlahSupplier int = 0
	var nextID int = 1
	var jumlahTransaksi int = 0
	var nextIDTransaksi int = 1
	
	fmt.Println("============================================")
	fmt.Println("    BangunIn - Sistem Informasi Supplier    ")
	fmt.Println("============================================")
	
	var jumlahProv int = 0
	isiDataLokasi(&jumlahProv)

	isiDummySup(&jumlahSupplier, &nextID)
	isiDummyTransaksi(&jumlahTransaksi, &nextIDTransaksi)

	for {
		fmt.Println("\n================ MENU UTAMA ================")
		fmt.Println("  1. Tampilkan Semua Data Supplier")
		fmt.Println("  2. Tambah Data Supplier")
		fmt.Println("  3. Ubah Data Supplier")
		fmt.Println("  4. Hapus Data Supplier")
		fmt.Println("  5. Catat Transaksi Baru")
		fmt.Println("  6. Tampilkan Riwayat Transaksi")
		fmt.Println("  7. Cari Data Supplier")
		fmt.Println("  8. Urutkan Data Supplier berdasarkan Rating")
		fmt.Println("  9. Statistik Supplier")
		fmt.Println("  10. Lihat dan Tambah Data Lokasi")
		fmt.Println("  0. Keluar")
		fmt.Println("============================================")
		fmt.Print("Pilihan Anda: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 0:
			fmt.Println("\nTerima kasih telah menggunakan BangunIn. Sampai jumpa!")
			return
		case 1:
			tampilkanSemuaSupplier(jumlahSupplier)
		case 2:
			tambahSupplier(&jumlahSupplier, &nextID, &jumlahProv)
		case 3:
			ubahSupplier(jumlahSupplier, &jumlahProv)
		case 4:
			hapusSupplier(&jumlahSupplier)
		case 5:
			catatTransaksi(jumlahSupplier, &jumlahTransaksi, &nextIDTransaksi)
		case 6:
			tampilkanRiwayatTransaksi(jumlahTransaksi)
		case 7:
			cariSupplier(jumlahSupplier, jumlahProv)
		case 8:
			urutkanSupplier(jumlahSupplier)
		case 9:
			tampilkanStatistik(jumlahSupplier)
		case 10:
			mode := "atur"
			aturAtauPilihLokasi(&jumlahProv, mode)
		default:
			fmt.Println("\nPilihan tidak valid. Silakan coba lagi.")
		}
	}
}