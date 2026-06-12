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

func main() {
	var p perangkat
	var totalPerangkat int
	seedData(&p, &totalPerangkat)
	tabel(p, totalPerangkat)
}
func tabel(p perangkat, totalPerangkat int) {
	var i int
	fmt.Printf("====================================================================================================\n")
	fmt.Printf("| %-3s | %-15s | %-15s | %-4s | %-12s | %-25s |\n", "No", "Nama", "Lokasi", "Daya (Watt)", "Durasi (Jam)", "Total Konsumsi (Watt/jam)")
	fmt.Printf("=--------------------------------------------------------------------------------------------------=\n")
	for i = 0; i < totalPerangkat; i++ {
		fmt.Printf("| %-3d | %-15s | %-15s | %-11d | %-12d | %-25d |\n", i+1, p[i].nama, p[i].lokasi, p[i].daya, p[i].durasi, p[i].konsumsiPerangkat)
	}
	fmt.Printf("=--------------------------------------------------------------------------------------------------=\n")
	fmt.Printf("| %-3d | %-15s | %-15s | %-11s | %-12s | %-25s |\n", totalPerangkat, "", "", "Watt", "W/Jam", "Total")
	fmt.Printf("====================================================================================================\n")
}
