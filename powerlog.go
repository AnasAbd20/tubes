package main

import "fmt"

const NMAX int = 100

type elektronik struct {
	nama, lokasi      string
	daya              int // Watt/jam
	durasi            int // Jam/hari
	konsumsiPerangkat int // Watt/hari
}
type perangkat [NMAX]elektronik

func seedData(p *perangkat, totalPerangkat *int) {
	// Fungsi untuk mengisi data awal perangkat elektronik ke dalam array daftar elektronik
	p[0] = elektronik{"TV", "Ruang_Tamu", 100, 4, 400}
	p[1] = elektronik{"Kulkas", "Dapur", 150, 24, 3600}
	p[2] = elektronik{"AC", "Kamar1", 200, 8, 1600}
	p[3] = elektronik{"Laptop", "Meja_Kerja", 50, 6, 300}
	p[4] = elektronik{"Lampu", "Ruang_Tamu", 20, 5, 100}
	p[5] = elektronik{"Kipas", "Kamar1", 40, 10, 400}
	p[6] = elektronik{"Setrika", "RuangCuci", 80, 2, 160}
	p[7] = elektronik{"TV", "RuangCuci", 120, 3, 360}
	p[8] = elektronik{"Kulkas", "RuangCuci", 180, 24, 4320}
	p[9] = elektronik{"AC", "Kamar2", 250, 6, 1500}
	*totalPerangkat = 10
}

func cetakLogo() {
	fmt.Println("==========================================================================")
	fmt.Println("██████╗  ██████╗ ██╗    ██╗███████╗██████╗     ██╗      ██████╗  ██████╗ ")
	fmt.Println("██╔══██╗██╔═══██╗██║    ██║██╔════╝██╔══██╗    ██║     ██╔═══██╗██╔════╝ ")
	fmt.Println("██████╔╝██║   ██║██║ █╗ ██║█████╗  ██████╔╝    ██║     ██║   ██║██║  ███╗")
	fmt.Println("██╔═══╝ ██║   ██║██║███╗██║██╔══╝  ██╔══██╗    ██║     ██║   ██║██║   ██║")
	fmt.Println("██║     ╚██████╔╝╚███╔███╔╝███████╗██║  ██║    ███████╗╚██████╔╝╚██████╔╝")
	fmt.Println("╚═╝      ╚═════╝  ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝    ╚══════╝ ╚═════╝  ╚═════╝ ")
}

func main() {
	var p perangkat
	var totalPerangkat int
	var done bool
	var pilihMenu int
	var urut string = "none"
	cetakLogo()
	seedData(&p, &totalPerangkat)
	for done = false; !done; {
		// Mencetak menu pilihan
		cetakMenu()
		fmt.Scan(&pilihMenu)
		// Menangani pilihan menu yang dipilih oleh pengguna
		switch pilihMenu {
		case 1:
			addPerangkat(&p, &totalPerangkat)
			// Setelah menambahkan perangkat, variabel string urut none karena array tidak terurut
			urut = "none"
		case 2:
			hapusPerangkat(&p, &totalPerangkat)
			urut = "none"
		case 3:
			editPerangkat(&p, totalPerangkat)
			// Setelah mengedit perangkat, variabel string urut none karena array tidak terurut
			urut = "none"
		case 4:
			// Cetak daftar perangkat elektronik beserta informasi lengkapnya
			cetakDaftarPerangkat(p, totalPerangkat)
		case 5:
			// Mengurutkan daftar perangkat elektronik
			sortingPerangkat(&p, totalPerangkat, &urut)
		case 6:
			// Mencari perangkat elektronik berdasarkan kategori yang dipilih dan mencetak informasi perangkat yang ditemukan
			searchingPerangkat(p, totalPerangkat, urut)
		case 7:
			// Mencetak statistik penggunaan energi dari semua perangkat elektronik
			cetakStatistik(p, totalPerangkat)
		case 0:
			done = true
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}

func cetakMenu() {
	// Fungsi untuk mencetak menu pilihan
	fmt.Println("==========================================================================")
	fmt.Println("Menu:")
	fmt.Println("1. Tambah Perangkat Elektronik")
	fmt.Println("2. Hapus Perangkat Elektronik")
	fmt.Println("3. Edit Perangkat Elektronik")
	fmt.Println("4. Tampilkan Daftar Perangkat Elektronik")
	fmt.Println("5. Urutkan Perangkat Elektronik")
	fmt.Println("6. Pencarian Perangkat Elektronik")
	fmt.Println("7. Tampilkan Statistik Penggunaan Energi")
	fmt.Println("0. Keluar")
	fmt.Println("==========================================================================")
	fmt.Print("Pilih menu: ")
}

func cetakKategori() {
	// Fungsi untuk mencetak kategori pada setiap perangkat elektronik
	fmt.Println("Kategori Perangkat Elektronik:")
	fmt.Println("1. Nama Perangkat")
	fmt.Println("2. Lokasi Perangkat")
	fmt.Println("3. Daya Perangkat")
	fmt.Println("4. Durasi Penggunaan")
}

func addPerangkat(p *perangkat, totalPerangkat *int) {
	// Fungsi untuk menambahkan perangkat elektronik ke dalam array daftar elektronik
	var n, i int
	fmt.Print("Masukkan jumlah perangkat yang ingin ditambahkan:")
	fmt.Scan(&n)
	// Cek apakah jumlah perangkat yang akan ditambahkan melebihi batas maksimum
	if *totalPerangkat+n > NMAX {
		fmt.Printf("Jumlah perangkat melebihi batas maksimum (%d)\n", NMAX)
		return
	}
	for i = *totalPerangkat; i < *totalPerangkat+n; i++ {
		fmt.Print("Masukkan nama perangkat: ")
		fmt.Scan(&p[i].nama)
		fmt.Print("Masukkan lokasi perangkat (gunakan underscore '_' untuk spasi): ")
		fmt.Scan(&p[i].lokasi)
		fmt.Print("Masukkan daya perangkat (Watt): ")
		fmt.Scan(&p[i].daya)
		fmt.Print("Masukkan durasi penggunaan perangkat (Jam/hari): ")
		fmt.Scan(&p[i].durasi)
		p[i].konsumsiPerangkat = hitungKonsumsiPerangkat(*p, i)
	}
	// Update total perangkat setelah penambahan
	*totalPerangkat += n
}

func hapusPerangkat(p *perangkat, totalPerangkat *int) {
	// Fungsi untuk menghapus perangkat elektronik dari array daftar elektronik
	var target string
	var selPer, i, found int
	fmt.Print("Masukkan nama perangkat yang ingin dihapus: ")
	fmt.Scan(&target)
	// Mencari indeks perangkat berdasarkan nama
	found = cariHitungNamaPerangkat(*p, *totalPerangkat, target)
	if found < 1 {
		fmt.Println("Perangkat tidak ditemukan")
		return
	} else if found > 1 {
		fmt.Print("Pilih nomor perangkat yang ingin dihapus: ")
		fmt.Scan(&selPer)
		selPer -= 1
	} else if found == 1 {
		selPer = cariNamaPerangkat(*p, *totalPerangkat, target)
	}
	for i = selPer; i < *totalPerangkat-1; i++ {
		p[i] = p[i+1]
	}
	// Update total perangkat setelah penghapusan
	*totalPerangkat--
}

func editPerangkat(p *perangkat, totalPerangkat int) {
	// Fungsi untuk mengedit informasi perangkat elektronik dalam array daftar elektronik
	var target string
	var selKat, found, selPer int
	fmt.Print("Masukkan nama perangkat yang ingin diedit: ")
	fmt.Scan(&target)
	// Mencari indeks perangkat berdasarkan nama
	found = cariHitungNamaPerangkat(*p, totalPerangkat, target)
	if found < 1 {
		fmt.Println("Perangkat tidak ditemukan")
		return
	} else if found == 1 {
		selPer = cariNamaPerangkat(*p, totalPerangkat, target)
	} else {
		fmt.Print("Pilih nomor perangkat yang ingin diedit: ")
		fmt.Scan(&selPer)
		selPer -= 1
	}
	cetakKategori()
	fmt.Print("Pilih kategori yang ingin diedit: ")
	fmt.Scan(&selKat)
	switch selKat {
	case 1:
		fmt.Print("Masukkan nama perangkat baru: ")
		fmt.Scan(&p[selPer].nama)
	case 2:
		fmt.Print("Masukkan lokasi perangkat baru: ")
		fmt.Scan(&p[selPer].lokasi)
	case 3:
		fmt.Print("Masukkan daya perangkat baru (Watt): ")
		fmt.Scan(&p[selPer].daya)
	case 4:
		fmt.Print("Masukkan durasi penggunaan perangkat baru (Jam/hari): ")
		fmt.Scan(&p[selPer].durasi)
	default:
		fmt.Println("Kategori tidak valid. Pengeditan dibatalkan.")
	}
}

func cariHitungNamaPerangkat(p perangkat, totalPerangkat int, target string) int {
	// Mencari target nama dalam data lalu dihitung jumlah yang sesuai nya
	var i int
	var found int
	var arrIndexFound [NMAX]int
	for i = 0; i < totalPerangkat; i++ {
		if strUpper(target) == strUpper(p[i].nama) {
			arrIndexFound[i] = 1
			found++
		}
	}
	if found < 1 {
		return found
	} else if found > 1 {
		cetakNamaFound(p, totalPerangkat, target, arrIndexFound)
	}
	return found

}

func cetakNamaFound(p perangkat, totalPerangkat int, target string, arrIndexFound [NMAX]int) {
	// Fungsi untuk mencetak daftar perangkat elektronik yang namanya sesuai dengan target beserta informasi lengkapnya
	var i int
	fmt.Printf("Daftar Perangkat Elektronik dengan Nama  %s:\n", strUpper(target))
	fmt.Printf("====================================================================================================\n")
	fmt.Printf("| %-3s | %-15s | %-15s | %-4s | %-12s | %-25s |\n", "No", "Nama", "Lokasi", "Daya (Watt)", "Durasi (Jam)", "Total Konsumsi (Watt/jam)")
	fmt.Printf("=--------------------------------------------------------------------------------------------------=\n")
	// Loop untuk mencetak
	for i = 0; i < totalPerangkat; i++ {
		// Pengecekan apakah index ini ada (1) dalam array arrIndexFound untuk dicetak
		if arrIndexFound[i] == 1 {
			fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", i+1, p[i].nama, p[i].lokasi, p[i].daya, p[i].durasi, p[i].konsumsiPerangkat)
		}
	}
	fmt.Printf("====================================================================================================\n")
}

func cetakDaftarPerangkat(p perangkat, totalPerangkat int) {
	// Fungsi untuk mencetak daftar perangkat elektronik beserta informasi lengkapnya
	var i int
	fmt.Println("Daftar Perangkat Elektronik:")
	fmt.Printf("====================================================================================================\n")
	fmt.Printf("| %-3s | %-15s | %-15s | %-4s | %-12s | %-25s |\n", "No", "Nama", "Lokasi", "Daya (Watt)", "Durasi (Jam)", "Total Konsumsi (Watt/jam)")
	fmt.Printf("=--------------------------------------------------------------------------------------------------=\n")
	for i = 0; i < totalPerangkat; i++ {
		fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", i+1, p[i].nama, p[i].lokasi, p[i].daya, p[i].durasi, p[i].konsumsiPerangkat)
	}
	fmt.Printf("=--------------------------------------------------------------------------------------------------=\n")
	fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", totalPerangkat, "", "", hitungTotalDaya(p, totalPerangkat), hitungTotalDurasi(p, totalPerangkat), hitungTotalKonsumsi(p, totalPerangkat))
	fmt.Printf("====================================================================================================\n")
}

func sortingPerangkat(p *perangkat, totalPerangkat int, urut *string) {
	// Fungsi untuk menu mengurutkan perangkat
	var selKat, selUrut int
	cetakKategori()
	fmt.Print("Pilih kategori yang ingin diurutkan: ")
	fmt.Scan(&selKat)
	for selKat < 1 || selKat > 4 {
		// Validasi input kategori, jika tidak valid, minta input ulang
		fmt.Println("Kategori tidak valid, silakan pilih kategori antara 1-4")
		fmt.Print("Pilih kategori yang ingin diurutkan: ")
		fmt.Scan(&selKat)
	}
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Ascending (A-Z atau 0-9)")
	fmt.Println("2. Descending (Z-A atau 9-0)")
	fmt.Print("Pilih urutan: ")
	fmt.Scan(&selUrut)
	for selUrut < 1 || selUrut > 2 {
		// Validasi input urutan, jika tidak valid, minta input ulang
		fmt.Println("Urutan tidak valid, silakan pilih urutan antara 1-2")
		fmt.Print("Pilih urutan: ")
		fmt.Scan(&selUrut)
	}
	switch selUrut {
	case 1:
		*urut = "naik"
		sortAscending(p, totalPerangkat, selKat)
	case 2:
		*urut = "turun"
		sortDescending(p, totalPerangkat, selKat)
	default:
		fmt.Println("Urutan tidak valid, silakan pilih urutan antara 1-2")
	}
}

func searchingPerangkat(p perangkat, totalPerangkat int, urut string) {
	// Fungsi untuk menu pencarian perangkat
	var selKat, selCetak int
	var target string
	cetakKategori()
	fmt.Print("Pilih kategori yang ingin dicari: ")
	fmt.Scan(&selKat)
	fmt.Print("Masukkan nilai yang ingin dicari: ")
	fmt.Scan(&target)
	fmt.Println("--------------------------------------")
	fmt.Println("1. Cetak informasi lengkap perangkat yang ditemukan (Mencetak semua perangkat yang sesuai dengan target)")
	fmt.Println("2. Cetak hanya nama dan lokasi perangkat yang ditemukan (Mencetak satu perangkat yang sesuai dengan target)")
	fmt.Println("--------------------------------------")
	fmt.Printf("Pilih jenis pencetakan hasil pencarian:")
	fmt.Scan(&selCetak)
	for selCetak < 1 || selCetak > 2 {
		// Validasi input pilihan cetak, jika tidak valid, minta input ulang
		fmt.Println("Pilihan tidak valid, silakan pilih antara 1-2")
		fmt.Print("Apakah Anda ingin mencetak informasi lengkap perangkat yang ditemukan?: ")
		fmt.Scan(&selCetak)
	}
	switch selCetak {
	case 1:
		// Mencetak informasi lengkap perangkat yang ditemukan menggunakan sequential search agar mencari semua perangkat yang sesuai dengan target
		cariPerangkatSeq(p, totalPerangkat, selKat, target)
	case 2:
		// Mencetak hanya nama dan lokasi perangkat yang ditemukan menggunakan binary search, sehingga hanya mencari satu perangkat yang sesuai dengan target
		if urut == "naik" {
			cariPerangkatBinAsc(p, totalPerangkat, selKat, target)
		} else if urut == "turun" {
			cariPerangkatBinDesc(p, totalPerangkat, selKat, target)
		} else {
			sortAscending(&p,totalPerangkat,1)
            fmt.Println("Data belum diurutkan, data akan otomatis di sorting untuk sementara dengan kategori nama secara ascending.")
            cariPerangkatBinAsc(p, totalPerangkat, selKat, target)
			
		}
	default:
		fmt.Println("Pilihan tidak valid, silakan pilih antara 1-2")
	}

}

func cariPerangkatSeq(p perangkat, totalPerangkat int, selKat int, target string) {
	// Fungsi untuk mencari perangkat elektronik berdasarkan kategori yang dipilih dan mencetak informasi lengkap perangkat yang ditemukan menggunakan sequential search, sehingga mencari semua perangkat yang sesuai dengan target
	var targetInt, idx int
	fmt.Printf("====================================================================================================\n")
	fmt.Printf("| %-3s | %-15s | %-15s | %-4s | %-12s | %-25s |\n", "No", "Nama", "Lokasi", "Daya (Watt)", "Durasi (Jam)", "Total Konsumsi (Watt/jam)")
	switch selKat {
	case 1:
		// Mencari berdasarkan nama perangkat
		target = strUpper(target) // Mengubah target nama menjadi huruf kapital
		for idx < totalPerangkat {
			// Mengubah nama perangkat menjadi huruf kapital
			if strUpper(p[idx].nama) == target {
				fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", idx+1, p[idx].nama, p[idx].lokasi, p[idx].daya, p[idx].durasi, p[idx].konsumsiPerangkat)
			}
			idx++
		}
	case 2:
		// Mencari berdasarkan lokasi perangkat
		target = strUpper(target) // Mengubah target lokasi menjadi huruf kapital
		for idx < totalPerangkat {
			// Mengubah lokasi perangkat menjadi huruf kapital
			if strUpper(p[idx].lokasi) == target {
				fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", idx+1, p[idx].nama, p[idx].lokasi, p[idx].daya, p[idx].durasi, p[idx].konsumsiPerangkat)
			}
			idx++
		}
	case 3:
		// Mencari berdasarkan daya perangkat
		targetInt = strToInt(target) // Mengubah target daya menjadi integer
		for idx < totalPerangkat {
			if p[idx].daya == targetInt {
				fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", idx+1, p[idx].nama, p[idx].lokasi, p[idx].daya, p[idx].durasi, p[idx].konsumsiPerangkat)
			}
			idx++
		}
	case 4:
		// Mencari berdasarkan durasi penggunaan perangkat
		targetInt = strToInt(target) // Mengubah target durasi menjadi integer
		for idx < totalPerangkat {
			if p[idx].durasi == targetInt {
				fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", idx+1, p[idx].nama, p[idx].lokasi, p[idx].daya, p[idx].durasi, p[idx].konsumsiPerangkat)
			}
			idx++
		}
	}
}

func cariPerangkatBinAsc(p perangkat, totalPerangkat int, selKat int, target string) {
	// Fungsi untuk mencari perangkat elektronik berdasarkan kategori yang dipilih dan mencetak hanya nama dan lokasi perangkat yang ditemukan menggunakan binary search, sehingga mencari satu perangkat yang sesuai dengan target, dengan asumsi data sudah diurutkan secara ascending
	var targetInt, left, mid, right int
	left = 0
	right = totalPerangkat - 1
	switch selKat {
	case 1:
		// Mencari berdasarkan nama perangkat
		target = strUpper(target) // Mengubah target nama menjadi huruf kapital
		for left <= right {
			mid = (left + right) / 2
			if strUpper(p[mid].nama) == target {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if strUpper(p[mid].nama) < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	case 2:
		// Mencari berdasarkan lokasi perangkat
		target = strUpper(target) // Mengubah target lokasi menjadi huruf kapital
		for left <= right {
			mid = (left + right) / 2
			if strUpper(p[mid].lokasi) == target {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if strUpper(p[mid].lokasi) < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	case 3:
		// Mencari berdasarkan daya perangkat
		targetInt = strToInt(target) // Mengubah target daya menjadi integer
		for left <= right {
			mid = (left + right) / 2
			if p[mid].daya == targetInt {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if p[mid].daya < targetInt {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	case 4:
		// Mencari berdasarkan durasi penggunaan perangkat
		targetInt = strToInt(target) // Mengubah target durasi menjadi integer
		for left <= right {
			mid = (left + right) / 2
			if p[mid].durasi == targetInt {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if p[mid].durasi < targetInt {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
}

func cariPerangkatBinDesc(p perangkat, totalPerangkat int, selKat int, target string) {
	// Fungsi untuk mencari perangkat elektronik berdasarkan kategori yang dipilih dan mencetak hanya nama dan lokasi perangkat yang ditemukan menggunakan binary search, sehingga mencari satu perangkat yang sesuai dengan target, dengan asumsi data sudah diurutkan secara descending
	var targetInt, left, mid, right int
	left = 0
	right = totalPerangkat - 1
	switch selKat {
	case 1:
		// Mencari berdasarkan nama perangkat
		target = strUpper(target) // Mengubah target nama menjadi huruf kapital
		for left <= right {
			mid = (left + right) / 2
			if strUpper(p[mid].nama) == target {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if strUpper(p[mid].nama) > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	case 2:
		// Mencari berdasarkan lokasi perangkat
		target = strUpper(target) // Mengubah target lokasi menjadi huruf kapital
		for left <= right {
			mid = (left + right) / 2
			if strUpper(p[mid].lokasi) == target {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if strUpper(p[mid].lokasi) > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	case 3:
		// Mencari berdasarkan daya perangkat
		targetInt = strToInt(target) // Mengubah target daya menjadi integer
		for left <= right {
			mid = (left + right) / 2
			if p[mid].daya == targetInt {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if p[mid].daya > targetInt {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	case 4:
		// Mencari berdasarkan durasi penggunaan perangkat
		targetInt = strToInt(target) // Mengubah target durasi menjadi integer
		for left <= right {
			mid = (left + right) / 2
			if p[mid].durasi == targetInt {
				fmt.Println("Data ditemukan")
				fmt.Printf("%d. Nama: %s, Lokasi: %s\n", mid+1, p[mid].nama, p[mid].lokasi)
				return
			} else if p[mid].durasi > targetInt {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
}

func sortAscending(p *perangkat, totalPerangkat int, selKat int) {
	// Fungsi untuk mengurutkan perangkat menggunakan insertion sort secara ascending berdasarkan kategori yang dipilih
	var i, pass int
	var temp elektronik
	pass = 1
	for pass < totalPerangkat {
		i = pass
		temp = p[i]
		switch selKat {
		case 1:
			// Mengurutkan berdasarkan nama perangkat
			for i > 0 && temp.nama < p[i-1].nama {
				p[i] = p[i-1]
				i = i - 1
			}
		case 2:
			// Mengurutkan berdasarkan lokasi perangkat
			for i > 0 && temp.lokasi < p[i-1].lokasi {
				p[i] = p[i-1]
				i = i - 1
			}
		case 3:
			// Mengurutkan berdasarkan daya perangkat
			for i > 0 && temp.daya < p[i-1].daya {
				p[i] = p[i-1]
				i = i - 1
			}
		case 4:
			// Mengurutkan berdasarkan durasi penggunaan perangkat
			for i > 0 && temp.durasi < p[i-1].durasi {
				p[i] = p[i-1]
				i = i - 1
			}
		}
		p[i] = temp
		pass++
	}
}

func sortDescending(p *perangkat, totalPerangkat int, selKat int) {
	// Fungsi untuk mengurutkan perangkat menggunakan selection sort secara descending berdasarkan kategori yang dipilih
	var i, pass, idx int
	var temp elektronik
	pass = 1
	for pass <= totalPerangkat-1 {
		idx = pass - 1
		i = pass + 1
		for i < totalPerangkat {
			switch selKat {
			case 1:
				// Mengurutkan berdasarkan nama perangkat
				if p[i].nama > p[idx].nama {
					idx = i
				}
			case 2:
				// Mengurutkan berdasarkan lokasi perangkat
				if p[i].lokasi > p[idx].lokasi {
					idx = i
				}
			case 3:
				// Mengurutkan berdasarkan daya perangkat
				if p[i].daya > p[idx].daya {
					idx = i
				}
			case 4:
				// Mengurutkan berdasarkan durasi penggunaan perangkat
				if p[i].durasi > p[idx].durasi {
					idx = i
				}
			}
			i++
		}
		temp = p[pass-1]
		p[pass-1] = p[idx]
		p[idx] = temp
		pass++
	}
}

func cariNamaPerangkat(p perangkat, totalPerangkat int, target string) int {
	// Fungsi untuk mencari perangkat elektronik berdasarkan nama dan mengembalikan indeksnya
	var i int
	var nama string
	target = strUpper(target) // Mengubah target nama menjadi huruf kapital
	for i < totalPerangkat {
		// Mengubah nama perangkat menjadi huruf kapital
		nama = strUpper(p[i].nama)
		if nama == target {
			return i
		}
		i++
	}
	return -1
}

func cariLokasiPerangkat(p perangkat, totalPerangkat int, target string) int {
	// Fungsi untuk mencari perangkat elektronik berdasarkan lokasi dan mengembalikan indeksnya
	var i int
	var lokasi string
	target = strUpper(target) // Mengubah target lokasi menjadi huruf kapital
	for i < totalPerangkat {
		// Mengubah lokasi perangkat menjadi huruf kapital
		lokasi = strUpper(p[i].lokasi)
		if lokasi == target {
			return i
		}
		i++
	}
	return -1
}

func strToInt(s string) int {
	// Fungsi untuk mengubah string menjadi integer, jika string tidak dapat diubah menjadi integer.
	var result int
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i]-'0') // Mengubah karakter angka menjadi integer dengan mengalikan hasil sebelumnya dengan 10 dan menambahkan nilai angka saat ini
		}
	}
	return result
}

func strUpper(s string) string {
	// Fungsi untuk mengubah string menjadi huruf kapital
	var result string
	var i int
	for i = 0; i < len(s); i++ {
		// Cek apakah karakter adalah huruf kecil (a-z)
		if s[i] >= 'a' && s[i] <= 'z' {
			result += string(s[i] - ('a' - 'A')) // Mengubah huruf kecil menjadi huruf kapital dengan mengurangi selisih antara 'a' dan 'A'
		} else {
			result += string(s[i]) // Jika karakter bukan huruf kecil, tetap tambahkan ke hasil tanpa perubahan
		}
	}
	return result
}

func hitungTotalDaya(p perangkat, totalPerangkat int) int {
	// Fungsi untuk menghitung total daya yang digunakan oleh semua perangkat elektronik per hari
	var totalDaya int
	var i int
	for i = 0; i < totalPerangkat; i++ {
		totalDaya += p[i].daya
	}
	return totalDaya
}
func hitungTotalDurasi(p perangkat, totalPerangkat int) int {
	// Fungsi untuk menghitung total durasi yang digunakan oleh semua perangkat dalam sehari
	var totalDurasi int
	var i int
	for i = 0; i < totalPerangkat; i++ {
		totalDurasi += p[i].durasi
	}
	return totalDurasi
}

func hitungKonsumsiPerangkat(p perangkat, idx int) int {
	// Fungsi untuk menghitung konsumsi energi yang digunakan oleh setiap perangkat elektronik per hari
	var konsumsi int
	konsumsi = p[idx].daya * p[idx].durasi
	return konsumsi
}

func hitungTotalKonsumsi(p perangkat, totalPerangkat int) int {
	// Fungsi untuk menghitung total konsumsi energi yang digunakan oleh semua perangkat elektronik per hari
	var totalKonsumsi int
	var i int
	for i = 0; i < totalPerangkat; i++ {
		totalKonsumsi += p[i].daya * p[i].durasi
	}
	return totalKonsumsi
}

func cetakStatistik(p perangkat, totalPerangkat int) {
	// Fungsi untuk mencetak statistik penggunaan energi dari semua perangkat elektronik
	var totalDaya, totalKonsumsi int
	var selkat int
	fmt.Println("Kategori untuk melihat statistik penggunaan energi:")
	fmt.Println("1. Nama Perangkat")
	fmt.Println("2. Lokasi Perangkat")
	fmt.Println("3. Total Daya dan Konsumsi Energi semua perangkat")
	fmt.Println("4. Grafik Konsumsi Energi")
	fmt.Print("Pilih kategori untuk melihat statistik penggunaan energi: ")
	fmt.Scan(&selkat)
	switch selkat {
	case 1:
		// Mencetak statistik penggunaan energi berdasarkan nama perangkat
		cetakStatistikKategori(p, totalPerangkat, selkat)
	case 2:
		// Mencetak statistik penggunaan energi berdasarkan lokasi perangkat
		cetakStatistikKategori(p, totalPerangkat, selkat)
	case 3:
		// Mencetak total daya dan konsumsi energi dari semua perangkat elektronik
		totalDaya = hitungTotalDaya(p, totalPerangkat)
		totalKonsumsi = hitungTotalKonsumsi(p, totalPerangkat)
		fmt.Printf("Total Daya yang Digunakan: %d Watt\n", totalDaya)
		fmt.Printf("Total Konsumsi Energi: %d Watt/hari\n", totalKonsumsi)
		fmt.Printf("Rata-rata Konsumsi Energi per Perangkat: %d Watt/hari\n", totalKonsumsi/totalPerangkat)
		fmt.Printf("Perangkat dengan Konsumsi Energi Tertinggi: %s (Konsumsi: %d Watt/hari)\n", p[cariPerangkatKonsumsiTertinggi(p, totalPerangkat)].nama, p[cariPerangkatKonsumsiTertinggi(p, totalPerangkat)].konsumsiPerangkat)
		fmt.Printf("Perangkat dengan Konsumsi Energi Terendah: %s (Konsumsi: %d Watt/hari)\n", p[cariPerangkatKonsumsiTerendah(p, totalPerangkat)].nama, p[cariPerangkatKonsumsiTerendah(p, totalPerangkat)].konsumsiPerangkat)
	case 4:
		// Mencetak statistik penggunaan energi berdasarkan kategori yang dipilih (nama atau lokasi) yang sama, sehingga dapat melihat statistik penggunaan energi dari setiap kelompok perangkat elektronik yang memiliki kategori yang sama
		fmt.Print("Pilih kategori (1: Nama, 2: Lokasi): ")
		fmt.Scan(&selkat)
		cetakGrafikKonsumsi(p, totalPerangkat, selkat)
		cetakGrafikUnit(p, totalPerangkat, selkat)
	default:
		fmt.Println("Kategori tidak valid, silakan pilih kategori antara 1-3")
	}
}
func cariPerangkatKonsumsiTertinggi(p perangkat, totalPerangkat int) int {
	// Fungsi untuk mencari perangkat dengan konsumsi energi tertinggi
	var idx, i int
	idx = 0
	for i = 1; i < totalPerangkat; i++ {
		if p[i].konsumsiPerangkat > p[idx].konsumsiPerangkat {
			idx = i
		}
	}
	return idx
}
func cariPerangkatKonsumsiTerendah(p perangkat, totalPerangkat int) int {
	// Fungsi untuk mencari perangkat dengan konsumsi energi terendah
	var idx, i int
	idx = 0
	for i = 1; i < totalPerangkat; i++ {
		if p[i].konsumsiPerangkat < p[idx].konsumsiPerangkat {
			idx = i
		}
	}
	return idx
}

func cetakStatistikKategori(p perangkat, totalPerangkat int, selkat int) {
	// Fungsi untuk mencetak statistik penggunaan energi berdasarkan kategori perangkat
	// Fungsi ini akan mencetak total daya dan konsumsi energi dari tiap perangkat elektronik yang dikelompokkan berdasarkan kategori yang dipilih (nama atau lokasi) yang sama, sehingga dapat melihat statistik penggunaan energi dari setiap kelompok perangkat elektronik yang memiliki kategori yang sama
	var i, j, k int
	var no int
	var foundPerangkat bool
	var oldValue, newValue, nama string
	var totalunit, totalDaya, totalKonsumsi int
	if selkat == 1 {
		nama = "Nama Perangkat"
	} else {
		nama = "Lokasi Perangkat"
	}
	fmt.Printf("==================================================================================\n")
	fmt.Println("Statistik penggunaan energi berdasarkan kategori perangkat:")
	fmt.Printf("==================================================================================\n")
	fmt.Printf("| %-3s | %-17s | %-4s | %-17s | %-25s |\n", "No", nama, "Unit", "Total Daya (Watt)", "Total Konsumsi (Watt/jam)")
	fmt.Printf("=--------------------------------------------------------------------------------=\n")
	for i = 0; i < totalPerangkat; i++ {
		if selkat == 1 {
			newValue = strUpper(p[i].nama) // Mengubah nama perangkat menjadi huruf kapital untuk memastikan perbandingan yang konsisten
		} else if selkat == 2 {
			newValue = strUpper(p[i].lokasi) // Mengubah lokasi perangkat menjadi huruf kapital untuk memastikan perbandingan yang konsisten
		}
		foundPerangkat = false
		for j = 0; j < i; j++ {
			if selkat == 1 {
				oldValue = strUpper(p[j].nama) // Mengubah nama perangkat menjadi huruf kapital untuk memastikan perbandingan yang konsisten
			} else if selkat == 2 {
				oldValue = strUpper(p[j].lokasi) // Mengubah lokasi perangkat menjadi huruf kapital untuk memastikan perbandingan yang konsisten
			}
			if newValue == oldValue {
				foundPerangkat = true
			}
		}
		if !foundPerangkat {
			// Inisialisasi
			totalunit = 0
			totalDaya = 0
			totalKonsumsi = 0
			// Loop untuk menghitung total unit, daya, dan konsumsi sesuai value (nama/lokasi) masing-masing
			for k = 0; k < totalPerangkat; k++ {
				if selkat == 1 {
					if strUpper(p[k].nama) == newValue {
						totalunit++
						totalDaya += p[k].daya
						totalKonsumsi += p[k].konsumsiPerangkat
					}
				} else if selkat == 2 {
					if strUpper(p[k].lokasi) == newValue {
						totalunit++
						totalDaya += p[k].daya
						totalKonsumsi += p[k].konsumsiPerangkat
					}
				}
			}
			if selkat == 1 {
				no++
				fmt.Printf("| %-3d | %-17s | %-4d | %-17d | %-25d |\n", no, newValue, totalunit, totalDaya, totalKonsumsi)
			} else if selkat == 2 {
				no++
				fmt.Printf("| %-3d | %-17s | %-4d | %-17d | %-25d |\n", no, newValue, totalunit, totalDaya, totalKonsumsi)
			}
		}
	}
	fmt.Printf("==================================================================================\n")
}

func cetakGrafikKonsumsi(p perangkat, totalPerangkat int, selkat int) {
	var i, j, k int
	var foundPerangkat bool
	var oldValue, newValue string
	var totalKonsumsi, skala, jumlahBalok int
	var b int

	fmt.Println("\n==================================================================")
	fmt.Println("Grafik Konsumsi Energi (dalam Watt/hari):")
	fmt.Println("1 balok (█) = 200 Watt/hari")
	fmt.Println("==================================================================")

	for i = 0; i < totalPerangkat; i++ {
		if selkat == 1 {
			newValue = strUpper(p[i].nama)
		} else if selkat == 2 {
			newValue = strUpper(p[i].lokasi)
		}

		foundPerangkat = false
		for j = 0; j < i; j++ {
			if selkat == 1 {
				oldValue = strUpper(p[j].nama)
			} else if selkat == 2 {
				oldValue = strUpper(p[j].lokasi)
			}
			if newValue == oldValue {
				foundPerangkat = true
			}
		}

		if !foundPerangkat {
			totalKonsumsi = 0
			for k = 0; k < totalPerangkat; k++ {
				if selkat == 1 {
					if strUpper(p[k].nama) == newValue {
						totalKonsumsi += p[k].konsumsiPerangkat
					}
				} else if selkat == 2 {
					if strUpper(p[k].lokasi) == newValue {
						totalKonsumsi += p[k].konsumsiPerangkat
					}
				}
			}
			// Skala: misal 1 balok mewakili 200 Watt/hari agar grafik tidak terlalu panjang ke kanan
			skala = 200
			jumlahBalok = totalKonsumsi / skala

			fmt.Printf("%-15s | ", newValue)

			// Cetak balok grafik menggunakan perulangan
			for b = 0; b < jumlahBalok; b++ {
				fmt.Print("█")
			}
			// Tampilkan nilai aslinya di ujung grafik
			fmt.Printf(" (%d Watt/hari)\n", totalKonsumsi)
		}
	}
	fmt.Println("==================================================================")
}
func cetakGrafikUnit(p perangkat, totalPerangkat int, selkat int) {
	var i, j, k int
	var foundPerangkat bool
	var oldValue, newValue string
	var totalUnit, skala, jumlahBalok int
	var b int

	fmt.Println("\n==================================================================")
	fmt.Println("Grafik Jumlah Unit Perangkat:")
	fmt.Println("1 Bintang (*) = 1 Unit")
	fmt.Println("==================================================================")

	for i = 0; i < totalPerangkat; i++ {
		if selkat == 1 {
			newValue = strUpper(p[i].nama)
		} else if selkat == 2 {
			newValue = strUpper(p[i].lokasi)
		}

		foundPerangkat = false
		for j = 0; j < i; j++ {
			if selkat == 1 {
				oldValue = strUpper(p[j].nama)
			} else if selkat == 2 {
				oldValue = strUpper(p[j].lokasi)
			}
			if newValue == oldValue {
				foundPerangkat = true
			}
		}

		if !foundPerangkat {
			totalUnit = 0
			for k = 0; k < totalPerangkat; k++ {
				if selkat == 1 {
					if strUpper(p[k].nama) == newValue {
						totalUnit++
					}
				} else if selkat == 2 {
					if strUpper(p[k].lokasi) == newValue {
						totalUnit++
					}
				}
			}
			// Skala: misal 1 balok mewakili 200 Watt/hari agar grafik tidak terlalu panjang ke kanan
			skala = 1
			jumlahBalok = totalUnit / skala
			fmt.Printf("%-15s | ", newValue)

			// Cetak balok grafik menggunakan perulangan
			for b = 0; b < jumlahBalok; b++ {
				fmt.Print("*")
			}
			// Tampilkan nilai aslinya di ujung grafik
			fmt.Printf(" (%d Unit)\n", totalUnit)
		}
	}
	fmt.Println("==================================================================")
}
