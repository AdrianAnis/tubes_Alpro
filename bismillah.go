package main

import "fmt"

const NMAX int = 100000

type dataTimTournament struct {
	namaTim   string
	jadwalTim string
	win       int
	lose      int
	skorTim   int
}

type dataTour [NMAX]dataTimTournament

var jumTimTour int
var dataTim dataTour

var timA, timB, JadwalMatch string
var skorA, skorB int

//========================================================================================================================================================================//

func main() {
	for {
		var pilihOpsiPeran int
		MenuUtama(&pilihOpsiPeran)

		switch pilihOpsiPeran {
		case 1:
			menuPanitia()
		case 2:
			menuPeserta()
		case 3:
			fmt.Println("\nTerima kasih telah menggunakan sistem!")
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}

// MENU SYSTEM
// //========================================================================================================================================================================//
func MenuUtama(pilih *int) {
	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t║   SELAMAT DATANG DI E-SPORTS TOURNAMENT TELKIF!  ║\t\n")
	fmt.Print("\t║                  Dibuat Oleh:                    ║\t\n")
	fmt.Print("\t║        Adrian Anis Pratama (103012400098)        ║\t\n")
	fmt.Print("\t║        Agiel Ghazi Abizar (103012400037)         ║\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t║==================================================║\t\n")
	fmt.Print("\t║ Silahkan Pilih Peran Anda Dibawah                ║\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t║ 1. Panitia Tournament                            ║\t\n")
	fmt.Print("\t║ 2. Pemain/Peserta Tournament                     ║\t\n")
	fmt.Print("\t║ 3. Keluar Aplikasi                               ║\t\n")
	fmt.Print("\t║                                    ╔═════════════║\t\n")
	fmt.Print("\t║                                    ║HALAMAN UTAMA║\t\n")
	fmt.Print("\t╚════════════════════════════════════╩═════════════╝\t\n")
	fmt.Print("\tMasukkan Pilihan Anda: ")
	fmt.Scan(pilih)
	fmt.Println()
}

func menuPanitia() {
	for {
		var pilihan int
		fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║          SELAMAT DATANG SEBAGAI PANITIA          ║\t\n")
		fmt.Print("\t║          DI E-SPORTS TOURNAMENT TELKIF!          ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║==================================================║\t\n")
		fmt.Print("\t║ Silahkan Pilih Opsi Yang Ada Dibawah             ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║ 1. Tambah TIM                                    ║\t\n")
		fmt.Print("\t║ 2. Mengubah/Mengedit Data TIM                    ║\t\n")
		fmt.Print("\t║ 3. Menghapus Data TIM                            ║\t\n")
		fmt.Print("\t║ 4. Mencari TIM                                   ║\t\n")
		fmt.Print("\t║ 5. Masukan Hasil Pertandingan                    ║\t\n")
		fmt.Print("\t║ 6. Tampilkan Klasemen                            ║\t\n")
		fmt.Print("\t║ 7. Halaman Utama                                 ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
		fmt.Print("\tMasukkan Pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahTim()
		case 2:
			ubahTim()
		case 3:
			hapusTim()
		case 4:
			cariTim()
		case 5:
			inputHasilPertandingan()
		case 6:
			tampilkanKlasemen()
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuPeserta() {
	for {
		var pilihan int
		fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║          SELAMAT DATANG SEBAGAI PEMAIN           ║\t\n")
		fmt.Print("\t║          DI E-SPORTS TOURNAMENT TELKIF!          ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║==================================================║\t\n")
		fmt.Print("\t║ Silahkan Pilih Opsi Yang Ada Dibawah             ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║ 1. Mencari TIM                                   ║\t\n")
		fmt.Print("\t║ 2. Tampilkan Klasemen                            ║\t\n")
		fmt.Print("\t║ 3. Halaman Utama                                 ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
		fmt.Print("\tMasukkan Pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			cariTim()
		case 2:
			tampilkanKlasemen()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

//========================================================================================================================================================================//
func tambahTim() {
	var i int
	var n int

	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║      Berapa total TIM yang akan Ditambahkan?     ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
	fmt.Print("\tMasukkan Pilihan Anda: ")
	fmt.Scan(&n)

	fmt.Print("\n\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║    Masukkan Nama Tim yang akan berpartisipasi    ║\t\n")
	fmt.Print("\t║    dan Jadwal pertandingan Dari TIM tersebut     ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")

	for i = 0; i < n; i++ {
		fmt.Printf("\n\tTim ke-%d:\n", i+1)
		fmt.Print("\tNama Tim: ")
		fmt.Scan(&dataTim[jumTimTour].namaTim)
		fmt.Print("\tJadwal Pertandingan: ")
		fmt.Scan(&dataTim[jumTimTour].jadwalTim)
		jumTimTour++
	}
	fmt.Println("\tData tim berhasil ditambahkan!")
}

func inputHasilPertandingan() {
	var idxA, idxB int

	fmt.Print("\n\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║      Masukkan Nama Tim yang akan bertanding      ║\t\n")
	fmt.Print("\t║         dan Jadwal pertandingan Hari ini         ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")

	fmt.Print("\tTanggal Pertandingan: ")
	fmt.Scan(&JadwalMatch)
	fmt.Print("\tTim A: ")
	fmt.Scan(&timA)
	fmt.Print("\tTim B: ")
	fmt.Scan(&timB)

	idxA = cariTimByNama(timA)
	idxB = cariTimByNama(timB)

	if idxA == -1 || idxB == -1 {
		fmt.Println("\tError: Tim tidak terdaftar!")
		return
	}

	fmt.Print("\tSkor Tim A: ")
	fmt.Scan(&skorA)
	fmt.Print("\tSkor Tim B: ")
	fmt.Scan(&skorB)

	if skorA > skorB {
		dataTim[idxA].win++
		dataTim[idxA].skorTim += 3
		dataTim[idxB].lose++
	} else if skorB > skorA {
		dataTim[idxB].win++
		dataTim[idxB].skorTim += 3
		dataTim[idxA].lose++
	}
	fmt.Println("\tHasil pertandingan berhasil disimpan!")
}

func ubahTim() {
	var nama string
	var idx int

	fmt.Print("\n\tMasukkan nama tim yang akan diubah: ")
	fmt.Scan(&nama)
	idx = cariTimByNama(nama)

	if idx == -1 {
		fmt.Println("\n\tTim tidak ditemukan!")
		return
	}

	fmt.Println("\n\tData saat ini:")
	fmt.Println("\t╔════════════════════════════════════════════╗")
	fmt.Printf("\t║ %-11s ║ %-6s ║ %4s ║ %4s ║ %5s ║\n", "TIM", "JADWAL", "W", "L", "SKOR")
	fmt.Println("\t╠═════════════╬════════╬══════╬══════╬═══════╣")
	fmt.Printf("\t║ %-11s ║ %-6s ║ %4d ║ %4d ║ %5d ║\n",
		dataTim[idx].namaTim,
		dataTim[idx].jadwalTim,
		dataTim[idx].win,
		dataTim[idx].lose,
		dataTim[idx].skorTim)
	fmt.Println("\t╚═════════════╩════════╩══════╩══════╩═══════╝")

	fmt.Print("\n\tMasukkan nama baru: ")
	fmt.Scan(&dataTim[idx].namaTim)
	fmt.Print("\tMasukkan jadwal baru: ")
	fmt.Scan(&dataTim[idx].jadwalTim)
	fmt.Println("\tData berhasil diupdate!")
}

func hapusTim() {
	var idx, i int
	var nama string

	fmt.Print("\n\tMasukkan nama tim yang akan dihapus: ")
	fmt.Scan(&nama)
	idx = cariTimByNama(nama)

	if idx == -1 {
		fmt.Println("\tTim tidak ditemukan!")
		return
	}

	for i = idx; i < jumTimTour-1; i++ {
		dataTim[i] = dataTim[i+1]
	}
	jumTimTour--
	fmt.Println("\tTim berhasil dihapus!")
}

func cariTim() {
	var nama string
	var idxSeq, idxBin int
	var sortedData dataTour

	fmt.Print("\n\tMasukkan nama tim: ")
	fmt.Scan(&nama)

	idxSeq = cariTimByNama(nama)
	sortedData = dataTim
	sortByNama(&sortedData, jumTimTour)
	idxBin = binarySearch(sortedData, jumTimTour, nama)

	fmt.Println("\n\t╔══════════════ HASIL PENCARIAN ═════════════╗")
	fmt.Println("\t║              Sequential Search             ║")
	fmt.Println("\t╠════════════════════════════════════════════╣")
	if idxSeq != -1 {
		fmt.Printf("\t║ %-15s: %-25s ║\n", "Nama Tim", dataTim[idxSeq].namaTim)
		fmt.Printf("\t║ %-15s: %-25s ║\n", "Jadwal", dataTim[idxSeq].jadwalTim)
		fmt.Printf("\t║ %-15s: %-25d ║\n", "Kemenangan", dataTim[idxSeq].win)
		fmt.Printf("\t║ %-15s: %-25d ║\n", "Kekalahan", dataTim[idxSeq].lose)
		fmt.Printf("\t║ %-15s: %-25d ║\n", "Skor", dataTim[idxSeq].skorTim)
	} else {
		fmt.Println("\t║            Tim tidak ditemukan!            ║")
	}
	fmt.Println("\t╠════════════════════════════════════════════╣")
	fmt.Println("\t║               Binary Search                ║")
	fmt.Println("\t╠════════════════════════════════════════════╣")
	if idxBin != -1 {
		fmt.Printf("\t║ %-15s: %-25s ║\n", "Nama Tim", sortedData[idxBin].namaTim)
		fmt.Printf("\t║ %-15s: %-25s ║\n", "Jadwal", sortedData[idxBin].jadwalTim)
		fmt.Printf("\t║ %-15s: %-25d ║\n", "Kemenangan", sortedData[idxBin].win)
		fmt.Printf("\t║ %-15s: %-25d ║\n", "Kekalahan", sortedData[idxBin].lose)
		fmt.Printf("\t║ %-15s: %-25d ║\n", "Skor", sortedData[idxBin].skorTim)
	} else {
		fmt.Println("\t║            Tim tidak ditemukan!            ║")
	}
	fmt.Println("\t╚════════════════════════════════════════════╝")
}

func tampilkanKlasemen() {
	var i int
	var tempSelection dataTour
	var tempInsertion dataTour

	if jumTimTour == 0 {
		fmt.Println("\n\tBelum ada data tim!")
		return
	}

	tempSelection = dataTim
	selectionSort(&tempSelection, jumTimTour)

	fmt.Println("\n\t╔══════════════════════════════════════════════╗")
	fmt.Println("\t║        KLASEMEN TOURNAMENT (SELECTION)       ║")
	fmt.Println("\t╠══════════════╦════════════╦═════╦═════╦══════╣")

	fmt.Printf("\t║   %-10s ║   %-9s║  %-3s║  %-3s║ %-5s║\n",
		"  TIM", " JADWAL", "W", "L", "SKOR")
	fmt.Println("\t╠══════════════╬════════════╬═════╬═════╬══════╣")

	for i = 0; i < jumTimTour; i++ {
		fmt.Printf("\t║     %-8s ║ %-6s ║%3d  ║%3d  ║%5d ║\n",
			tempSelection[i].namaTim,
			tempSelection[i].jadwalTim,
			tempSelection[i].win,
			tempSelection[i].lose,
			tempSelection[i].skorTim)
	}
	fmt.Println("\t╚══════════════╩════════════╩═════╩═════╩══════╝")

	tempInsertion = dataTim
	insertionSort(&tempInsertion, jumTimTour)

	fmt.Println("\n\t╔══════════════════════════════════════════════╗")
	fmt.Println("\t║        KLASEMEN TOURNAMENT (INSERTION)       ║")
	fmt.Println("\t╠══════════════╦════════════╦═════╦═════╦══════╣")

	fmt.Printf("\t║   %-10s ║   %-9s║  %-3s║  %-3s║ %-5s║\n",
		"  TIM", " JADWAL", "W", "L", "SKOR")
	fmt.Println("\t╠══════════════╬════════════╬═════╬═════╬══════╣")

	for i = 0; i < jumTimTour; i++ {
		fmt.Printf("\t║     %-8s ║ %-6s ║%3d  ║%3d  ║%5d ║\n",
			tempInsertion[i].namaTim,
			tempInsertion[i].jadwalTim,
			tempInsertion[i].win,
			tempInsertion[i].lose,
			tempInsertion[i].skorTim)
	}
	fmt.Println("\t╚══════════════╩════════════╩═════╩═════╩══════╝")
}

//========================================================================================================================================================================//

func cariTimByNama(nama string) int {
	var i int

	for i = 0; i < jumTimTour; i++ {
		if dataTim[i].namaTim == nama {
			return i
		}
	}
	return -1
}

func binarySearch(data dataTour, n int, target string) int {
	var low, high, mid int

	low = 0
	high = n - 1

	for low <= high {
		mid = (low + high) / 2

		if data[mid].namaTim == target {
			return mid
		} else if data[mid].namaTim < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func sortByNama(data *dataTour, n int) {
	var i,j int
	var k dataTimTournament

	for i = 1; i < n; i++ {
		k = data[i]
		j = i - 1
		for j >= 0 && data[j].namaTim > k.namaTim {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = k
	}
}

func selectionSort(data *dataTour, n int) {
	var i,maxIdx,j int

	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if data[j].skorTim > data[maxIdx].skorTim {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
}

func insertionSort(data *dataTour, n int) {
	var i,j int
	var k dataTimTournament

	for i = 1; i < n; i++ {
		k = data[i]
		j = i - 1
		for j >= 0 && data[j].skorTim < k.skorTim {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = k
	}
}
