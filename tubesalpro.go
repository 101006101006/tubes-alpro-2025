package main

import (
	"fmt"
	"sort"
)

type NilaiUjian struct {
	UH, UTS, UAS, UKK float64
}

type Mapel struct {
	Nama  string
	Nilai NilaiUjian
}

type Siswa struct {
	NIS    string
	Nama   string
	Mapels []Mapel
}

type InfoMapel struct {
	Nama string
	Max  float64
	Min  float64
}

var dataSiswa = []Siswa{
	{"                                                                     001", "Ibra", []Mapel{
		{"Matematika", NilaiUjian{80, 85, 90, 88}},
		{"Bahasa Indonesia", NilaiUjian{78, 80, 82, 85}},
	}},
	{"                                                                     002", "Alex", []Mapel{
		{"Matematika", NilaiUjian{75, 80, 70, 72}},
		{"Bahasa Indonesia", NilaiUjian{85, 88, 90, 87}},
	}},
	{"                                                                     003", "Enca", []Mapel{
		{"Matematika", NilaiUjian{90, 92, 95, 94}},
		{"Bahasa Indonesia", NilaiUjian{80, 85, 88, 86}},
	}},
}

func menu() {
	fmt.Println("                                                  ╔═════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                                                  ║                    CREATED BY NEIZSHIA AND IBRA                     ║")
	fmt.Println("                                                  ╠═════════════════════════════════════════════════════════════════════╣")
	fmt.Println("                                                  ║                     PENDATAAN NILAI UJIAN SISWA                     ║")
	fmt.Println("                                                  ╠═════════════════════════════════════════════════════════════════════╣")
	fmt.Println("                                                  ╠═════════════════════════════════════════════════════════════════════╣")
	fmt.Println("                                                  ║                                                                     ║")
	fmt.Println("                                                  ║   1. Tampilkan Semua Data                                           ║")
	fmt.Println("                                                  ║   2. Tampilkan Rata-Rata Tertinggi Per Mapel                        ║")
	fmt.Println("                                                  ║   3. Tampilkan Nilai Maksimum & Minimum Per Mapel                   ║")
	fmt.Println("                                                  ║   4. Cari Data (Sequential Search berdasarkan Nama)                 ║")
	fmt.Println("                                                  ║   5. Cari Data (Binary Search berdasarkan NIS)                      ║")
	fmt.Println("                                                  ║   6. Urutkan Data (Selection Sort berdasarkan Rata-rata)            ║") 
	fmt.Println("                                                  ║   7. Urutkan Data (Insertion Sort berdasarkan Nama)                 ║")
	fmt.Println("                                                  ║   8. Tampilkan Data                                                 ║")
	fmt.Println("                                                  ║   0. Keluar                                                         ║")
	fmt.Println("                                                  ║                                                                     ║")
	fmt.Println("                                                  ║                                                                     ║")                                                                     
	fmt.Println("                                                  ╚═════════════════════════════════════════════════════════════════════╝")
  	fmt.Print("                                                   Pilih menu: ")
}

func rataRata(n NilaiUjian) float64 {
	return (n.UH + n.UTS + n.UAS + n.UKK) / 4
}

func rataRataSiswa(siswa Siswa) float64 {
	total := 0.0
	for i := 0; i < len(siswa.Mapels); i++ {
		total += rataRata(siswa.Mapels[i].Nilai)
	}
	return total / float64(len(siswa.Mapels))
}

func tampilSemuaData() {
	fmt.Println("                                                                        ╔══════════════════════════╗")
	fmt.Println("                                                                        ║=== DATA SELURUH SISWA ===║")
	fmt.Println("                                                                        ╚══════════════════════════╝")
	for i := 0; i < len(dataSiswa); i++ {
		siswa := dataSiswa[i]
		fmt.Printf("[%d] %s (%s)\n", i+1, siswa.Nama, siswa.NIS)
		for j := 0; j < len(siswa.Mapels); j++ {
			mapel := siswa.Mapels[j]
			fmt.Printf("                                                  - %s: UH=%.0f, UTS=%.0f, UAS=%.0f, UKK=%.0f\n",
				mapel.Nama, mapel.Nilai.UH, mapel.Nilai.UTS, mapel.Nilai.UAS, mapel.Nilai.UKK)
		}
		fmt.Printf("                                                      >> Rata-rata nilai: %.2f\n", rataRataSiswa(siswa))
		fmt.Println("                             ")
	}
}

func tampilRataRataTertinggiPerMapel() {
	fmt.Println("                                                              ╔══════════════════════════════════════════════╗")
	fmt.Println("                                                              ║=== RATA-RATA TERTINGGI PER MATA PELAJARAN ===║")
	fmt.Println("                                                              ╚══════════════════════════════════════════════╝")
	
	mapelTeratas := []Mapel{}
	nilaiTertinggi := []float64{}
	namaSiswa := []string{}

	for i := 0; i < len(dataSiswa); i++ {
		siswa := dataSiswa[i]
		for j := 0; j < len(siswa.Mapels); j++ {
			mapel := siswa.Mapels[j]
			rata := rataRata(mapel.Nilai)
			
			found := false
			for k := 0; k < len(mapelTeratas); k++ {
				if mapelTeratas[k].Nama == mapel.Nama {
					if rata > nilaiTertinggi[k] {
						nilaiTertinggi[k] = rata
						namaSiswa[k] = siswa.Nama
					}
					found = true
				}
			}

			if !found {
				mapelTeratas = append(mapelTeratas, mapel)
				nilaiTertinggi = append(nilaiTertinggi, rata)
				namaSiswa = append(namaSiswa, siswa.Nama)
			}
		}
	}

	for i := 0; i < len(mapelTeratas); i++ {
		fmt.Printf("                                                      - %s: %s (%.2f)\n", mapelTeratas[i].Nama, namaSiswa[i], nilaiTertinggi[i])
	}
}

func tampilNilaiMaxMin() {
	fmt.Println("                                                               ╔══════════════════════════════════════════╗")
	fmt.Println("                                                               ║=== NILAI MAKSIMUM & MINIMUM PER MAPEL ===║")
	fmt.Println("                                                               ╚══════════════════════════════════════════╝")
	
	var namaMapel []string
	var nilaiMax []float64
	var nilaiMin []float64

	for i := 0; i < len(dataSiswa); i++ {
		siswa := dataSiswa[i]
		for j := 0; j < len(siswa.Mapels); j++ {
			mapel := siswa.Mapels[j]
			rata := rataRata(mapel.Nilai)

			idx := -1
			found := false
			for k := 0; k < len(namaMapel); k++ {
				if namaMapel[k] == mapel.Nama {
					idx = k
					found = true
					// tidak break, tetap lanjut sampai akhir
				}
			}

			if !found {
				// mapel baru, tambahkan ke slice
				namaMapel = append(namaMapel, mapel.Nama)
				nilaiMax = append(nilaiMax, rata)
				nilaiMin = append(nilaiMin, rata)
			} else {
				// update max dan min jika perlu
				if rata > nilaiMax[idx] {
					nilaiMax[idx] = rata
				}
				if rata < nilaiMin[idx] {
					nilaiMin[idx] = rata
				}
			}
		}
	}
	for i := 0; i < len(namaMapel); i++ {
		fmt.Printf("                                                      - %s: Max = %.2f, Min = %.2f\n", namaMapel[i], nilaiMax[i], nilaiMin[i])
	}
}

func sequentialSearch(nama string) {
	fmt.Println("                                                               ╔═══════════════════════════════════╗")
	fmt.Println("                                                               ║=== CARI SISWA BERDASARKAN NAMA ===║")
	fmt.Println("                                                               ╚═══════════════════════════════════╝")
	
	status := false
	for i := 0; i < len(dataSiswa); i++ {
		if dataSiswa[i].Nama == nama {
			fmt.Printf("                                                  Ditemukan: %s (%s)\n", dataSiswa[i].Nama, dataSiswa[i].NIS)
			status = true
		}
	}
	if !status {
		fmt.Println("                                                  Tidak ditemukan.")
	}
}

func binarySearch(nis string) {
	fmt.Println("                                                                ╔══════════════════════════════════╗")
	fmt.Println("                                                                ║=== CARI SISWA BERDASARKAN NIS ===║")
	fmt.Println("                                                                ╚══════════════════════════════════╝")
	
	// Pastikan data sudah terurut berdasarkan NIS
	sort.Slice(dataSiswa, func(i, j int) bool {
		return dataSiswa[i].NIS < dataSiswa[j].NIS
	})

	low, high := 0, len(dataSiswa)-1
	for low <= high {
		mid := (low + high) / 2
		if dataSiswa[mid].NIS == nis {
			fmt.Printf("                                                  Ditemukan: %s (%s)\n", dataSiswa[mid].Nama, dataSiswa[mid].NIS)
			return
		} else if dataSiswa[mid].NIS < nis {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("                                                      Tidak ditemukan.")
}

func selectionSortByRata() {
	for i := 0; i < len(dataSiswa); i++ {
		minIdx := i
		for j := i + 1; j < len(dataSiswa); j++ {
			if rataRataSiswa(dataSiswa[j]) < rataRataSiswa(dataSiswa[minIdx]) {
				minIdx = j
			}
		}
		dataSiswa[i], dataSiswa[minIdx] = dataSiswa[minIdx], dataSiswa[i]
	}
	fmt.Println("                                                      Data diurutkan berdasarkan rata-rata (Selection Sort).")
}

func insertionSortByNama() {
	n := len(dataSiswa)
	for i := 1; i <= n-1; i++ {
		pass := dataSiswa[i]
		j := i - 1
		for j >= 0 && dataSiswa[j].Nama > pass.Nama {
			dataSiswa[j+1] = dataSiswa[j]
			j = j - 1
		}
		dataSiswa[j+1] = pass
	}
	fmt.Println("                                                      Data diurutkan berdasarkan nama (Insertion Sort).")
}

func tampilRekursif(index int) {
	if index >= len(dataSiswa) {
		return
	}
	siswa := dataSiswa[index]
	fmt.Printf("                                                          [%d] %s (%s)\n", index+1, siswa.Nama, siswa.NIS)
	for j := 0; j < len(siswa.Mapels); j++ {
		mapel := siswa.Mapels[j]
		fmt.Printf("  - %s: UH=%.0f, UTS=%.0f, UAS=%.0f, UKK=%.0f\n",
			mapel.Nama, mapel.Nilai.UH, mapel.Nilai.UTS, mapel.Nilai.UAS, mapel.Nilai.UKK)
	}
	fmt.Printf("                                                           >> Rata-rata: %.2f\n", rataRataSiswa(siswa))
	fmt.Println("                              ")
	tampilRekursif(index + 1)
}

func main() {
	var pilihan int
	for {
		menu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tampilSemuaData()
		case 2:
			tampilRataRataTertinggiPerMapel()
		case 3:
			tampilNilaiMaxMin()
		case 4:
			var nama string
			fmt.Print("                                                   Masukkan nama yang dicari: ")
			fmt.Scan(&nama)
			sequentialSearch(nama)
		case 5:
			var nis string
			fmt.Print("                                                   Masukkan NIS yang dicari: ")
			fmt.Scan(&nis)
			binarySearch(nis)
		case 6:
			selectionSortByRata()
			tampilSemuaData()
		case 7:
			insertionSortByNama()
			tampilSemuaData()
		case 8:
			fmt.Println("                                                             ╔══════════════════════════════╗")
			fmt.Println("                                                             ║=== TAMPIL DATA (REKURSIF) ===║")
			fmt.Println("                                                             ╚══════════════════════════════╝")
			tampilRekursif(0)
		case 0:
			fmt.Println("                                                  Terimakasih :)")
			return
		default:
			fmt.Println("                                                 Pilihan tidak valid.")
		}
	}
}