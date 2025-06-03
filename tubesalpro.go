package main
import(
	"fmt" //menampilkan teks ke layar
	"os" //menjalankan perintah sistem (misalnya clear screen)
	"os/exec"
	"runtime" //untuk mendeteksi OS
	//"sort"
)

const NMAX int = 100

type NilaiUjian struct {
	UH, UTS, UAS, UKK float64
}

type MataPelajaran struct {
	Nama 	string
	Nilai 	NilaiUjian
}

type Siswa struct {
	NIS		string
	Nama	string
	Mapel	MataPelajaran
}

//Struktur array utama untuk menyimpan semua data
type ArrData struct {
	TabSiswa [NMAX]Siswa
	TabMapel [NMAX]MataPelajaran
	TabUjian [NMAX]NilaiUjian
}

//membuat jarak spasi saat mencetak teks ke layar
var indent string = "                                 " 
var indentTab string = "                       "

func main() {
	var data ArrData
	var nData int = 4
	var pilihan int
	
	data.TabUjian[0] = NilaiUjian{UH: 85, UTS: 88, UAS: 90, UKK: 92}
	data.TabMapel[0] = MataPelajaran{Nama: "Matematika", Nilai: data.TabUjian[0]}
	data.TabSiswa[0] = Siswa{NIS: "00001", Nama: "Budi", Mapel: data.TabMapel[0]}
	
	data.TabUjian[1] = NilaiUjian{UH: 97, UTS: 98, UAS: 92, UKK: 94}
	data.TabMapel[1] = MataPelajaran{Nama: "Bahasa_Inggris", Nilai: data.TabUjian[1]}
	data.TabSiswa[1] = Siswa{NIS: "00001", Nama: "Budi", Mapel: data.TabMapel[1]}

	data.TabUjian[2] = NilaiUjian{UH: 78, UTS: 80, UAS: 82, UKK: 85}
	data.TabMapel[2] = MataPelajaran{Nama: "Bahasa_Indonesia", Nilai: data.TabUjian[2]}
	data.TabSiswa[2] = Siswa{NIS: "00002", Nama: "Ani", Mapel: data.TabMapel[2]}
	
	data.TabUjian[3] = NilaiUjian{UH: 67, UTS: 83, UAS: 90, UKK: 80}
	data.TabMapel[3] = MataPelajaran{Nama: "Matematika", Nilai: data.TabUjian[3]}
	data.TabSiswa[3] = Siswa{NIS: "00002", Nama: "Ani", Mapel: data.TabMapel[3]}
	
	for {
		menu()
		fmt.Scan(&pilihan)
		
		switch pilihan {
		case 1:
			inputDataBaru(&data, &nData)
		case 2:
			clearScreen()
			tampilSemuaData(data, nData)
		case 3:
			ubahDataTerdaftar(&data, nData)
		case 4:
			hapusDataByNama(&data, &nData)
		case 5:
			tampilRataRataSiswaPerMapel(data, nData)
		case 6:
			tampilNilaiMaksMinPerMapel(data, nData)
		case 7:
			clearScreen()
			kotakTemplate("           === CARI DATA DENGAN SEQUENTIAL SEARCH NAMA ===          ")
			var cariNama string
			fmt.Print(indent + "Masukkan nama siswa yang dicari: ")
			fmt.Scan(&cariNama)

			var hasil [NMAX]int
			var jumlah int

			hasil, jumlah = sequentialSearchByNama(data, nData, cariNama)

			if jumlah == 0 {
				fmt.Println(indent + "Data tidak ditemukan.")
			} else {
				fmt.Printf(indent + "Ditemukan %d data dengan nama %s:\n", jumlah, cariNama)
				for i := 0; i < jumlah; i++ {
					idx := hasil[i]
					fmt.Printf(indent + "- NIS: %s, Mapel: %s\n", data.TabSiswa[idx].NIS, data.TabSiswa[idx].Mapel.Nama)
				}
			}

			var pause string
			fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
			fmt.Scan(&pause)

		case 8:
			clearScreen()
			kotakTemplate("      === CARI DATA DENGAN BINARY SEARCH BERDASARKAN NIS ===         ")
			var target string
			fmt.Print(indent + "Cari NIS: ")
			fmt.Scan(&target)
			selectionSortByNIS(&data, nData)
			
			idx := binarySearchByNIS(data, nData, target)
			
			if idx != -1 {
				fmt.Printf(indent + "Ditemukan siswa dengan NIS %s:\n", target)
				fmt.Printf(indent + "Nama: %s, Mata Pelajaran: %s\n", data.TabSiswa[idx].Nama, data.TabSiswa[idx].Mapel.Nama)
			} else {
				fmt.Printf(indent + "Siswa dengan NIS %s tidak ditemukan.\n", target)
			}
	
			var pause string
			fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
			fmt.Scan(&pause)
		case 9:
			clearScreen()		
			kotakTemplate("         === DATA DIURUTKAN BERDASARKAN MATA PELAJARAN ===           ")
			selectionSortByMapel(&data, nData)
			tampilSemuaData(data, nData)
		case 10:
			clearScreen()
			kotakTemplate("             === DATA DIURUTKAN BERDASARKAN NAMA ===                 ")
			insertionSortByNama(&data, nData)
			tampilSemuaData(data, nData)
		case 0:
			clearScreen()
			kotakTemplate("                          Terimakasih :)                             ")
			return
		default:
			kotakTemplate("                      Pilihan tidak valid                            ")
		}
	}
}

func clearScreen() { //membersihkan layar terminal sesuai dengan OS yang digunakan
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func menu() {
	clearScreen() // Bersihkan layar sebelum menampilkan menu
	fmt.Println(indent + "╔═════════════════════════════════════════════════════════════════════╗")
	fmt.Println(indent + "║                    CREATED BY NEIZSHIA AND IBRA                     ║")
	fmt.Println(indent + "╠═════════════════════════════════════════════════════════════════════╣")
	fmt.Println(indent + "║                     PENDATAAN NILAI UJIAN SISWA                     ║")
	fmt.Println(indent + "╠═════════════════════════════════════════════════════════════════════╣")
	fmt.Println(indent + "╠═════════════════════════════════════════════════════════════════════╣")
	fmt.Println(indent + "║   1. Input Data Baru                                                ║")
	fmt.Println(indent + "║   2. Tampilkan Semua Data                                           ║")
	fmt.Println(indent + "║   3. Ubah Data Terdaftar                                            ║")
	fmt.Println(indent + "║   4. Hapus Data Terdaftar berdasarkan Nama                          ║")
	fmt.Println(indent + "║   5. Tampilkan Rata-Rata Per Mapel                                  ║")
	fmt.Println(indent + "║   6. Tampilkan Nilai Maksimum & Minimum Per Mapel                   ║")
	fmt.Println(indent + "║   7. Cari Data (Sequential Search berdasarkan Nama)                 ║")
	fmt.Println(indent + "║   8. Cari Data (Binary Search berdasarkan NIS)                      ║")
	fmt.Println(indent + "║   9. Urutkan Data (Selection Sort berdasarkan Mata pelajaran)       ║") 
	fmt.Println(indent + "║   10.Urutkan Data (Insertion Sort berdasarkan Nama)                 ║")
	fmt.Println(indent + "║   0. Keluar                                                         ║")
	fmt.Println(indent + "╚═════════════════════════════════════════════════════════════════════╝")
  	fmt.Print(indent + "Pilih menu: ")
}

// Prosedur untuk menampilkan teks melalui template kotak
func kotakTemplate(pesan string) {
	fmt.Println(indent + "╔═════════════════════════════════════════════════════════════════════╗")
	fmt.Printf(indent + "║%69s║\n", pesan)
	fmt.Println(indent + "╚═════════════════════════════════════════════════════════════════════╝")
}

func inputDataBaru(data *ArrData, n *int) {
	if *n >= NMAX {
		*n = NMAX
	}
	
	for {
		clearScreen()
		kotakTemplate("                       === INPUT DATA BARU ===                       ")
		fmt.Println()
		kotakTemplate("              *Untuk spasi mohon memakai (_) underscore              ")
		fmt.Print(indent + "Masukkan NIS: ")
		fmt.Scan(&data.TabSiswa[*n].NIS)
		fmt.Print(indent + "Masukkan Nama: ")
		fmt.Scan(&data.TabSiswa[*n].Nama)
		fmt.Print(indent + "Masukkan Mata Pelajaran: ")
		fmt.Scan(&data.TabSiswa[*n].Mapel.Nama)
		fmt.Print(indent + "Masukkan Nilai UH: ")
		fmt.Scan(&data.TabSiswa[*n].Mapel.Nilai.UH)
		fmt.Print(indent + "Masukkan Nilai UTS: ")
		fmt.Scan(&data.TabSiswa[*n].Mapel.Nilai.UTS)
		fmt.Print(indent + "Masukkan Nilai UAS: ")
		fmt.Scan(&data.TabSiswa[*n].Mapel.Nilai.UAS)
		fmt.Print(indent + "Masukkan Nilai UH: ")
		fmt.Scan(&data.TabSiswa[*n].Mapel.Nilai.UKK)
		
		*n = *n + 1
		
		var pause string
		fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
		fmt.Scan(&pause)
		return
	}
}

func tampilSemuaData(data ArrData, n int) {
	var no int = 1
	kotakTemplate("                      === DATA SELURUH SISWA ===                     ")
	
	fmt.Println(indentTab + "┌─────┬───────┬────────────┬──────────────────────┬────────┬────────┬────────┬────────┐")
	fmt.Printf(indentTab + "│ %3s │ %-5s │ %-10s │ %-20s │ %6s │ %6s │ %6s │ %6s │\n", "No.", "NIS", "Nama", "Mata Pelajaran", "UH", "UTS", "UAS", "UKK")
	fmt.Println(indentTab + "├─────┼───────┼────────────┼──────────────────────┼────────┼────────┼────────┼────────┤")	

	for i := 0; i < n; i++ {
		fmt.Printf(indentTab + "│ %3d │ %5s │ %-10s │ %-20s │ %6.2f │ %6.2f │ %6.2f │ %6.2f │\n", 
			no,
			data.TabSiswa[i].NIS, 
			data.TabSiswa[i].Nama, 
			data.TabSiswa[i].Mapel.Nama, 
			data.TabSiswa[i].Mapel.Nilai.UH, 
			data.TabSiswa[i].Mapel.Nilai.UTS, 
			data.TabSiswa[i].Mapel.Nilai.UAS, 
			data.TabSiswa[i].Mapel.Nilai.UKK)
		no++
	}
	
	fmt.Println(indentTab + "└─────┴───────┴────────────┴──────────────────────┴────────┴────────┴────────┴────────┘")
	var pause string
	fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
	fmt.Scan(&pause)
	return
}

func ubahDataTerdaftar(data *ArrData, n int) {
	var cariData string
	var hasilCari [NMAX]int
	var nCari, nUbah int
	var no int = 1
	var arrUbah Siswa
	clearScreen()
	kotakTemplate("                     === UBAH DATA TERDAFTAR ===                     ")
	fmt.Print(indent + "Cari nama siswa: ")
	fmt.Scan(&cariData)
	
	hasilCari, nCari = sequentialSearch(*data, n, cariData, "Nama")
	
	if nCari == 0 {
        fmt.Println(indent + "Data tidak ditemukan.")
		var pause string
		fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
		fmt.Scan(&pause)
		return
    }
	
	fmt.Println(indentTab + "┌─────┬───────┬────────────┬──────────────────────┬────────┬────────┬────────┬────────┐")
	fmt.Printf(indentTab + "│ %3s │ %-5s │ %-10s │ %-20s │ %6s │ %6s │ %6s │ %6s │\n", "No.", "NIS", "Nama", "Mata Pelajaran", "UH", "UTS", "UAS", "UKK")
	fmt.Println(indentTab + "├─────┼───────┼────────────┼──────────────────────┼────────┼────────┼────────┼────────┤")
	for i := 0; i < nCari; i++ {
			fmt.Printf(indentTab + "│ %3d │ %5s │ %-10s │ %-20s │ %6.2f │ %6.2f │ %6.2f │ %6.2f │\n",
			no,
			data.TabSiswa[hasilCari[i]].NIS, 
			data.TabSiswa[hasilCari[i]].Nama, 
			data.TabSiswa[hasilCari[i]].Mapel.Nama, 
			data.TabSiswa[hasilCari[i]].Mapel.Nilai.UH, 
			data.TabSiswa[hasilCari[i]].Mapel.Nilai.UTS, 
			data.TabSiswa[hasilCari[i]].Mapel.Nilai.UAS, 
			data.TabSiswa[hasilCari[i]].Mapel.Nilai.UKK)
		no++
	}
	fmt.Println(indentTab + "└─────┴───────┴────────────┴──────────────────────┴────────┴────────┴────────┴────────┘")
	fmt.Print(indent + "Pilih nomor data yang ingin diubah: ")
	fmt.Scan(&nUbah)
	nUbah = nUbah - 1
	
	if nUbah < 0 || nUbah >= nCari {
		fmt.Println("Pilihan nomor data tidak valid")
		return
	}

	fmt.Println()
	kotakTemplate("              *Untuk spasi mohon memakai (_) underscore              ")
	fmt.Print(indent + "Masukkan NIS baru: ")
	fmt.Scan(&arrUbah.NIS)
	fmt.Print(indent + "Masukkan Nama baru: ")
	fmt.Scan(&arrUbah.Nama)
	fmt.Print(indent + "Masukkan Mata Pelajaran baru: ")
	fmt.Scan(&arrUbah.Mapel.Nama)
	fmt.Print(indent + "Masukkan Nilai UH baru: ")
	fmt.Scan(&arrUbah.Mapel.Nilai.UH)
	fmt.Print(indent + "Masukkan Nilai UTS baru: ")
	fmt.Scan(&arrUbah.Mapel.Nilai.UTS)
	fmt.Print(indent + "Masukkan Nilai UAS baru: ")
	fmt.Scan(&arrUbah.Mapel.Nilai.UAS)
	fmt.Print(indent + "Masukkan Nilai UKK baru: ")
	fmt.Scan(&arrUbah.Mapel.Nilai.UKK)
	
	idx := hasilCari[nUbah]
    data.TabSiswa[idx].NIS = arrUbah.NIS
    data.TabSiswa[idx].Nama = arrUbah.Nama
    data.TabSiswa[idx].Mapel.Nama = arrUbah.Mapel.Nama
    data.TabSiswa[idx].Mapel.Nilai.UH = arrUbah.Mapel.Nilai.UH
    data.TabSiswa[idx].Mapel.Nilai.UTS = arrUbah.Mapel.Nilai.UTS
    data.TabSiswa[idx].Mapel.Nilai.UAS = arrUbah.Mapel.Nilai.UAS
    data.TabSiswa[idx].Mapel.Nilai.UKK = arrUbah.Mapel.Nilai.UKK
	
	var pause string
	fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
	fmt.Scan(&pause)
	return
}

func hapusDataByNama(data *ArrData, n *int) {
	var nama string
	var hasilCari [NMAX]int
	var nCari, pilihan int

	clearScreen()
	kotakTemplate("                    === HAPUS DATA SISWA ===                    ")
	fmt.Print(indent + "Masukkan nama siswa yang ingin dihapus: ")
	fmt.Scan(&nama)

	hasilCari, nCari = sequentialSearch(*data, *n, nama, "Nama")

	if nCari == 0 {
		fmt.Println(indent + "Data tidak ditemukan.")
		var pause string
		fmt.Print("\n" + indent + "Tekan ENTER untuk kembali... ")
		fmt.Scan(&pause)
		return
	}

	fmt.Println()
	fmt.Println(indentTab + "Data ditemukan:")
	for i := 0; i < nCari; i++ {
		idx := hasilCari[i]
		fmt.Printf(indentTab+"%d. %s - %s - %s\n", i+1, data.TabSiswa[idx].NIS, data.TabSiswa[idx].Nama, data.TabSiswa[idx].Mapel.Nama)
	}
	fmt.Print("\n" + indent + "Pilih nomor data yang ingin dihapus (1 - ", nCari, "): ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > nCari {
		fmt.Println(indent + "Pilihan tidak valid.")
		var pause string
		fmt.Print("\n" + indent + "Tekan ENTER untuk kembali... ")
		fmt.Scan(&pause)
		return
	}

	idxHapus := hasilCari[pilihan-1]

	// Geser semua elemen setelah idxHapus ke kiri
	for i := idxHapus; i < *n-1; i++ {
		data.TabSiswa[i] = data.TabSiswa[i+1]
	}

	// Kurangi jumlah data
	*n = *n - 1

	fmt.Println(indent + "Data berhasil dihapus.")
	var pause string
	fmt.Print("\n" + indent + "Tekan ENTER untuk kembali... ")
	fmt.Scan(&pause)
}


func sequentialSearch(data ArrData, n int, dicari, mode string) ([NMAX]int, int) {
	// var idx int = -1
	var idx [NMAX]int
	var j int = 0
	
	for i := 0; i < n; i++ {
		switch mode {
		case "Nama":
			if data.TabSiswa[i].Nama == dicari {
				idx[j] = i
				j++
			}
		case "NIS":
			if data.TabSiswa[i].NIS == dicari {
				idx[j] = i
				j++
			}
		case "Mapel":
			if data.TabSiswa[i].Mapel.Nama == dicari {
				idx[j] = i
				j++
			}
		default:
			// Mode tidak dikenali, skip
		}
	}
	
	return idx, j
}


func tampilRataRataSiswaPerMapel(data ArrData, n int) {
	var namaSiswa, namaMapel string
	clearScreen()
	kotakTemplate("            === RATA-RATA SISWA UNTUK MATA PELAJARAN ===           ")
	fmt.Print(indent + "Masukkan nama siswa: ")
	fmt.Scan(&namaSiswa)

	fmt.Print(indent + "Masukkan nama mata pelajaran: ")
	fmt.Scan(&namaMapel)

	rata := hitungRataRataMapelUntukSiswa(data, n, namaSiswa, namaMapel)

	if rata == 0 {
		fmt.Println(indent + "Data tidak ditemukan atau siswa belum mengambil mapel tersebut.")
		var pause string
		fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
		fmt.Scan(&pause)
		return
	} else {
		fmt.Printf(indent + "%s untuk rata-rata %s: %.2f\n", namaSiswa, namaMapel, rata)
	}
	
	var pause string
	fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
	fmt.Scan(&pause)
	return
}


func hitungRataRataMapelUntukSiswa(data ArrData, n int, namaSiswa, mapel string) float64 {
	var hasilCari [NMAX]int
	var nCari int
	var total float64 = 0
	var jumlah int = 0

	// Cari semua data siswa yang sesuai nama
	hasilCari, nCari = sequentialSearch(data, n, namaSiswa, "Nama")

	if nCari == 0 {
		return 0 // Jika yang dicari tidak ditemukan
	}

	for i := 0; i < nCari; i++ {
		idx := hasilCari[i]
		if data.TabSiswa[idx].Mapel.Nama == mapel {
			nilai := data.TabSiswa[idx].Mapel.Nilai
			rata := (nilai.UH + nilai.UTS + nilai.UAS + nilai.UKK) / 4
			total += rata
			jumlah++
		}
	}

	if jumlah == 0 {
		return 0 // Tidak ada mapel yang cocok
	}

	return total / float64(jumlah)
}


func tampilNilaiMaksMinPerMapel(data ArrData, n int) {
	var namaMapel string
	var hasilCari [NMAX]int
	var nCari int

	clearScreen()
	kotakTemplate("         === NILAI MAKSIMUM & MINIMUM PER MATA PELAJARAN ===        ")
	fmt.Print(indent + "Masukkan nama mata pelajaran yang ingin dicek: ")
	fmt.Scan(&namaMapel)

	hasilCari, nCari = sequentialSearch(data, n, namaMapel, "Mapel")

	if nCari == 0 {
		fmt.Printf("Mapel %s tidak ditemukan atau belum ada data.\n", namaMapel)
		var pause string
		fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
		fmt.Scan(&pause)
		return
	}

	idxMax := hasilCari[0]
	idxMin := hasilCari[0]
	nilaiMax := hitungRata(data.TabSiswa[idxMax].Mapel.Nilai)
	nilaiMin := nilaiMax

	for i := 1; i < nCari; i++ {
		idx := hasilCari[i]
		rata := hitungRata(data.TabSiswa[idx].Mapel.Nilai)

		if rata > nilaiMax {
			nilaiMax = rata
			idxMax = idx
		}
		if rata < nilaiMin {
			nilaiMin = rata
			idxMin = idx
		}
	}

	fmt.Printf(indent + "\nMapel dicek : %s\n", namaMapel)
	fmt.Printf(indent + "  Nilai Tertinggi : %s\n", data.TabSiswa[idxMax].Nama)
	fmt.Printf(indent + "  Nilai Terendah  : %s\n", data.TabSiswa[idxMin].Nama)
	
	var pause string
	fmt.Print("\n" + indent + "Ketik apapun lalu ENTER untuk kembali... ")
	fmt.Scan(&pause)
	return
}

func hitungRata(nilai NilaiUjian) float64 {
	return (nilai.UH + nilai.UTS + nilai.UAS + nilai.UKK) / 4
}

func binarySearchByNIS(data ArrData, n int, target string) int {
	low := 0
	high := n - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2
		if data.TabSiswa[mid].NIS == target {
			return mid // Ketemu, kembalikan index
		} else if data.TabSiswa[mid].NIS < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Tidak ketemu
}

func selectionSortByNIS(data *ArrData, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data.TabSiswa[j].NIS < data.TabSiswa[minIdx].NIS {
				minIdx = j
			}
		}
		if minIdx != i {
			data.TabSiswa[i], data.TabSiswa[minIdx] = data.TabSiswa[minIdx], data.TabSiswa[i]
		}
	}
}

func sequentialSearchByNama(data ArrData, n int, nama string) ([NMAX]int, int) {
	var hasil [NMAX]int
	var jumlah int = 0

	for i := 0; i < n; i++ {
		if data.TabSiswa[i].Nama == nama {
			hasil[jumlah] = i
			jumlah++
		}
	}
	return hasil, jumlah
}

func selectionSortByMapel(data *ArrData, n int) {
	var i, j, idxMin int
	var temp Siswa

	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if data.TabSiswa[j].Mapel.Nama < data.TabSiswa[idxMin].Mapel.Nama {
				idxMin = j
			}
		}
		if idxMin != i {
			temp = data.TabSiswa[i]
			data.TabSiswa[i] = data.TabSiswa[idxMin]
			data.TabSiswa[idxMin] = temp
		}
	}
}

func insertionSortByNama(data *ArrData, n int) {
	var i, j int
	var temp Siswa

	for i = 1; i < n; i++ {
		temp = data.TabSiswa[i]
		j = i
		for j > 0 && data.TabSiswa[j-1].Nama > temp.Nama {
			data.TabSiswa[j] = data.TabSiswa[j-1]
			j--
		}
		data.TabSiswa[j] = temp
	}
}
